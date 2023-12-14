package global

import (
	"fmt"
	"time"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func InitLogger() {
	logger, err := log.LoggerFromConfigAsString("<seelog>\n\t<outputs formatid=\"main\">\n\t\t<console />\n\t\t<!-- <file path=\"./logs/col-middleware.log\"/> -->\n\t\t<rollingfile type=\"size\" filename=\"logs/backend.log\" maxsize=\"500000000\" maxrolls=\"5\" archivepath=\"logs/backend.zip\" />\n\t</outputs>\n\t<formats>\n\t\t<format id=\"main\" format=\"%Date %Time [%Level] %Msg%n\"/>\n\t</formats>\n</seelog>\n")
	if err != nil {
		fmt.Println("parse seelog.xml error")
		return
	}

	log.ReplaceLogger(logger)

	defer log.Flush()
	log.Info("init Seelog!")
}

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Infof("[IN] %15s %s %s",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
		)
		c.Next()

		log.Infof("[OUT] %3d %13v %15s %s %s",
			c.Writer.Status(),
			time.Now().Sub(start),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}

func NewHttpLog() *HttpLog {
	return new(HttpLog)
}

type HttpLog struct {
}

func (self *HttpLog) SetPrefix(prefix string) {}

func (self *HttpLog) Printf(format string, v ...interface{}) {
	log.Debugf(format, v)
}

func (self *HttpLog) Println(v ...interface{}) {
	log.Debug(v)
}
