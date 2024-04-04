package ansi

import "fmt"

type Color int

const (
	Red Color = iota
	Green
	Yellow
	Blue
)

var (
	ANSIColors = map[Color]string{
		Red:    "\033[31m",
		Green:  "\033[32m",
		Yellow: "\033[33m",
		Blue:   "\033[34m",
	}
)

func (c Color) Wrap(format string, args ...interface{}) string {
	return ANSIColors[c] + fmt.Sprintf(format, args...) + "\033[0m"
}
