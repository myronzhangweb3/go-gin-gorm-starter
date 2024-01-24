package global

import (
	"fmt"
	"time"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func InitLogger() {
	// 定义JSON格式的日志配置
	logger, err := log.LoggerFromConfigAsString(`
<seelog>
	<outputs formatid="json">
		<console />
		<rollingfile type="size" filename="logs/backend.log" maxsize="500000000" maxrolls="5" archivepath="logs/backend.zip" />
	</outputs>
	<formats>
		<format id="json" format='{"time":"%Date %Time","level":"%LEV","message":"%Msg"}%n'/>
	</formats>
</seelog>
`)
	if err != nil {
		fmt.Println("parse seelog config error:", err)
		return
	}

	log.ReplaceLogger(logger)

	defer log.Flush()
	log.Info("init Seelog!")
}

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Infof(`{"action":"IN","client_ip":"%s","method":"%s","path":"%s"}`,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
		)
		c.Next()

		log.Infof(`{"action":"OUT","status":%d,"duration":"%v","client_ip":"%s","method":"%s","path":"%s"}`,
			c.Writer.Status(),
			time.Since(start),
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
	log.Debugf(format, v...)
}

func (self *HttpLog) Println(v ...interface{}) {
	log.Debug(v...)
}
