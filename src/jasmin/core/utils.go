package core

import (
	"fmt"
	"strings"
)

func joinBySpace(args ...interface{}) string {
	var builder strings.Builder

	for i, arg := range args {
		str := fmt.Sprint(arg)
		if len(str) > 0 && i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(str)
	}

	return builder.String()
}

func static(isStatic bool) string {
	if isStatic {
		return "static"
	}
	return ""
}
