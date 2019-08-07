package wafer

import (
	"fmt"
	"io"
	"strings"
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
	ls := splitToLines(s)
	width := maxWidth(ls)
	w.writeVertical(width)
	for _, l := range ls {
		w.write(width, l)
	}
	w.writeVertical(width)
}

func splitToLines(s string) []string {
	return strings.Split(s, "\n")
}

func maxWidth(ss []string) int {
	var max int
	for _, s := range ss {
		if current := len(s); max < current {
			max = current
		}
	}

	return max
}

func (w *Wafer) writeVertical(width int) {
	w.builder.WriteString(fmt.Sprintf("%s\n", verticalLine(width+4)))
}

func (w *Wafer) write(width int, s string) {
	s = fmt.Sprintf("%s%s", s, strings.Repeat(" ", width-len(s)))
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
