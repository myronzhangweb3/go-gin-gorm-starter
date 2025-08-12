package global

import (
	"go-gin-gorm-starter/utils/dbutil"
	"go-gin-gorm-starter/utils/logging"
)

func init() {
	logging.InitLogger("demo")
	dbutil.InitDB()
}
