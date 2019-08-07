package wafer

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		"one line": {
			input:    "aiueo",
			expected: "---------\n| aiueo |\n---------\n",
		},
		"multiple lines": {
			input:    "aiueo\naiueo",
			expected: "---------\n| aiueo |\n| aiueo |\n---------\n",
		},
		"multiple lines of different width": {
			input:    "aiueo\naiueoaiueo",
			expected: "--------------\n| aiueo      |\n| aiueoaiueo |\n--------------\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var b strings.Builder
			w := New(&b)
			w.Append(test.input)
			w.Render()
			actual := b.String()
			if actual != test.expected {
				t.Errorf("unexpected result rendered by Render: got %q, expect %q\n", actual, test.expected)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type input struct {
		width int
		s     string
	}
	tests := map[string]struct {
		input    input
		expected string
	}{
		"same width": {
			input: input{
				width: 5,
				s:     "aiueo",
			},
			expected: "aiueo",
		},
		"different width": {
			input: input{
				width: 3,
				s:     "aiueo",
			},
			expected: "aiu\neo",
		},
		"multiple lines in different width": {
			input: input{
				width: 3,
				s:     "aiueo\naiueo\naiueo",
			},
			expected: "aiu\neo\naiu\neo\naiu\neo",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := wrap(test.input.width, test.input.s)
			if actual != test.expected {
				t.Errorf("unexpected string wrapped by wrap: got %q, expect %q\n", actual, test.expected)
			}
		})
	}
}

func TestCalculateMaxWidth(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected int
	}{
		"single byte": {
			input:    []string{"aiueo"},
			expected: 5,
		},
		"multi byte": {
			input:    []string{"アイウエオ"},
			expected: 5,
		},
		"single byte in multi elements": {
			input:    []string{"aiueo", "aiueoaiueo"},
			expected: 10,
		},
		"multi byte in multi elements": {
			input:    []string{"アイウエオ", "あいうえおアイウエオ"},
			expected: 10,
		},
		"single and multi byte in multi elements": {
			input:    []string{"aiueo", "アイウエオ"},
			expected: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := calculateMaxWidth(test.input)
			if actual != test.expected {
				t.Errorf("unexpected max width returned by calculatedMaxWidth: got %d, expect %d\n", actual, test.expected)
			}
		})
	}
}
