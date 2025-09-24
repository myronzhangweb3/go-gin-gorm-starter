package time2

import (
	"runtime"
	"time"

	log "github.com/cihub/seelog"
)

// TimeConsume ...
func TimeConsume(start time.Time) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	log.Debug(funcName, " cost:", time.Since(start).String())
}
