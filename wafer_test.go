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
