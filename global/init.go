package global

import (
	"go-gin-gorm-starter/pkg/dbutil"
	"go-gin-gorm-starter/pkg/logging"
)

func init() {
	logging.InitLogger("demo")
	dbutil.InitDB()
}
