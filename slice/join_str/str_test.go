package join_str

import "testing"

var (
	elt1 = "zttvi-p"
	elt2 = "xibei"
	elt3 = "lt-jiangsu-fog_docker_fogcdn-svg-18"
	elt4 = "dx-neimenggu-baotou-sn6-172-31-253-216"
)

func BenchmarkADD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ADD(elt1, elt2, elt3, elt4)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(elt1, elt2, elt3, elt4)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Buffer(elt1, elt2, elt3, elt4)
	}
}

func BenchmarkBufferPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BufferPool(elt1, elt2, elt3, elt4)
	}
}
