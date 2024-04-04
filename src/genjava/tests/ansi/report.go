package ansi

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

func (c Color) Wrap(msg string) string {
	return ANSIColors[c] + msg + "\033[0m"
}
