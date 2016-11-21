package treesort

import "bytes"
import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)

	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree)String() string {
	var vals []int
	vals = appendValues(vals,t)
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _, val := range vals {
		if buf.Len() > len("[") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", val)
	}
	buf.WriteByte(']')
	return buf.String()
}
