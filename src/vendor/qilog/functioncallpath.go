package qilog

import (
	"fmt"
	"runtime"
	"strings"
)

func CallPath(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	_, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("%s: %d: ", chopPath(file), line)
}

func chopPath(original string) string {
	i := strings.LastIndex(original, "src/")
	if i == -1 {
		return original
	} else {
		return original[i+4:]
	}
}
