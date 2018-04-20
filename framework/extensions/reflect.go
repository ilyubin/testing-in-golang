package extensions

import (
	"runtime"
)

func FunctionName() string {
	return caller(2)
}

func CallerName() string {
	return caller(3)
}

func CallerName2() string {
	return caller(4)
}

func caller(i int) string {
	pc, _, _, ok := runtime.Caller(i)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}
	return ""
}
