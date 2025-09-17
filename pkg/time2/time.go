package time2

import (
	log "github.com/cihub/seelog"
	"runtime"
	"time"
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
