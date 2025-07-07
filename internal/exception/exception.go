package exception

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/sajad-dev/gingo/internal/config"
)

var DEBUG string = config.Config.DEBUG

func Callers() []string {
	pc := make([]uintptr, 50)
	runtimeCallerNum := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:runtimeCallerNum])

	var callers []string
	for {
		frame, more := frames.Next()
		callers = append(callers, fmt.Sprintf("Called from %s\n\t%s:%d\n", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}

	}
	return callers
}
func Exception (err error) string {
	callers := Callers()
	callersSplit := fmt.Sprintf("Error :( %s : Callers ===> %s",err,strings.Join(callers," -> "))
	color.Red("We have Error !!! :(")
	if DEBUG == "false"{
		log.Println(callersSplit)	
	}else {
		color.Red(callersSplit)
	}
	return callersSplit
}
