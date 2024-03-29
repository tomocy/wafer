package wafer

import (
	"fmt"
	"io"
	"strings"
)

const (
	maxWidth = 40
)

func New(dest io.Writer) *Wafer {
	return &Wafer{
		dest: dest,
	}
}

type Wafer struct {
	dest    io.Writer
	builder strings.Builder
}

func (w *Wafer) Append(s string) {
	wrapped := wrap(maxWidth, s)
	ls := splitToLines(wrapped)
	width := calculateMaxWidth(ls)
	w.writeVertical(width)
	for _, l := range ls {
		w.write(width, l)
	}
	w.writeVertical(width)
}

func wrap(width int, s string) string {
	var b strings.Builder
	var lineLen int
	for i, r := range s {
		lineLen++
		b.WriteRune(r)
		if r == '\n' {
			lineLen = 0
			continue
		}
		if i != len(s)-1 && width <= lineLen {
			lineLen = 0
			b.WriteRune('\n')
		}
	}

	return b.String()
}

func splitToLines(s string) []string {
	return strings.Split(s, "\n")
}

func calculateMaxWidth(ss []string) int {
	var max int
	for _, s := range ss {
		if current := len([]rune(s)); max < current {
			max = current
		}
	}

	return max
}

func (w *Wafer) writeVertical(width int) {
	w.builder.WriteString(fmt.Sprintf("%s\n", verticalLine(width+4)))
}

func (w *Wafer) write(width int, s string) {
	s = fmt.Sprintf("%s%s", s, strings.Repeat(" ", width-len([]rune(s))))
	w.builder.WriteString(fmt.Sprintf("| %s |\n", s))
}

func verticalLine(width int) string {
	var b strings.Builder
	b.WriteString(strings.Repeat("-", width))

	return b.String()
}

func (w *Wafer) Render() {
	fmt.Fprint(w.dest, w.builder.String())
}
