package bigint

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

type Int big.Int

func New(x *big.Int) *Int {
	return (*Int)(x)
}

func NewArray(ints []*big.Int) []*Int {
	result := make([]*Int, len(ints))
	for i, _ := range ints {
		result[i] = New(ints[i])
	}
	return result
}

func NewArrayFromInterface(ints interface{}) ([]*Int, error) {
	var result []*Int
	jsonStr, err := json.Marshal(ints)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonStr, &result)
	return result, err
}

func FromInterface(x interface{}) (*Int, error) {
	switch v := x.(type) {
	case int64:
		return FromInt64(v), nil
	case *int64:
		return FromInt64(*v), nil
	case string:
		return FromString(v)
	case *string:
		return FromString(*v)
	case *big.Int:
		return New(v), nil
	default:
		return nil, errors.New("unexpected type")
	}
}

func FromInt64(x int64) *Int {
	return new(Int).FromBigInt(big.NewInt(x))
}

func FromString(x string) (*Int, error) {
	if x == "" {
		return nil, nil
	}
	a := big.NewInt(0)
	b, ok := a.SetString(x, 10)

	if !ok {
		return nil, fmt.Errorf("cannot create Int from string")
	}

	return New(b), nil
}

func (b *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Int) UnmarshalJSON(text []byte) error {
	str := string(text[1 : len(text)-1])
	a := big.NewInt(0)
	_, ok := a.SetString(str, 10)
	if !ok {
		return fmt.Errorf("cannot create Int from string")
	}
	*b = *New(a)
	return nil
}

func (b *Int) Value() (driver.Value, error) {
	return (*big.Int)(b).String(), nil
}

func (b *Int) Scan(value interface{}) error {

	var i sql.NullString

	if err := i.Scan(value); err != nil {
		return err
	}

	if _, ok := (*big.Int)(b).SetString(i.String, 10); ok {
		return nil
	}

	return fmt.Errorf("Error converting type %T into Int", value)
}

func (b *Int) toBigInt() *big.Int {
	return (*big.Int)(b)
}

func (b *Int) Sub(x *Int) *Int {
	return (*Int)(big.NewInt(0).Sub(b.toBigInt(), x.toBigInt()))
}

func (b *Int) Add(x *Int) *Int {
	return (*Int)(big.NewInt(0).Add(b.toBigInt(), x.toBigInt()))
}

func (b *Int) Mul(x *Int) *Int {
	return (*Int)(big.NewInt(0).Mul(b.toBigInt(), x.toBigInt()))
}

func (b *Int) Div(x *Int) *Int {
	return (*Int)(big.NewInt(0).Div(b.toBigInt(), x.toBigInt()))
}

func (b *Int) Neg() *Int {
	return (*Int)(big.NewInt(0).Neg(b.toBigInt()))
}

func (b *Int) ToUInt64() uint64 {
	return b.toBigInt().Uint64()
}

func (b *Int) ToInt64() int64 {
	return b.toBigInt().Int64()
}

// same as New()
func (b *Int) FromBigInt(x *big.Int) *Int {
	return (*Int)(x)
}

func (b *Int) String() string {
	return b.toBigInt().String()
}
