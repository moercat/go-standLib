package join_str

import (
	"bytes"
	"strings"
	"sync"
)

func ADD(elt1, elt2, elt3, elt4 string) string {
	return elt1 + "::" + elt2 + "::" + elt3 + "::" + elt4
}

func Join(elt ...string) string {
	return strings.Join(elt, "::")
}

func Buffer(elt1, elt2, elt3, elt4 string) string {
	var buffer bytes.Buffer
	buffer.Grow(7)
	buffer.WriteString(elt1)
	buffer.WriteString("::")
	buffer.WriteString(elt2)
	buffer.WriteString("::")
	buffer.WriteString(elt3)
	buffer.WriteString("::")
	buffer.WriteString(elt4)
	return buffer.String()
}

func BufferPool(elt1, elt2, elt3, elt4 string) string {
	var buffer = bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer bufferPool.Put(buffer)
	buffer.WriteString(elt1)
	buffer.WriteString("::")
	buffer.WriteString(elt2)
	buffer.WriteString("::")
	buffer.WriteString(elt3)
	buffer.WriteString("::")
	buffer.WriteString(elt4)
	return buffer.String()
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
