package if_var

var (
	I int
	M = make(map[string]int)
)

func Number1() {

	if v, ok := M["aa"]; ok {
		I = v
	}
	for i := 0; i < 1000; i++ {
		I++
	}
}

func Number2() {

	v, ok := M["aa"]
	if ok {
		I = v
	}
	for i := 0; i < 1000; i++ {
		I++
	}
}

var last *SardineCoverStrategy

type SardineCoverStrategy struct {
	TTL                   int   `json:"ttl"`
	Timestamp             int64 `json:"timestamp"`
	ExpireTime            int64 `json:"-"`
	UseSNodeComputeWeight bool  `json:"use_s_node_compute_weight"`
	UseAvailablePlan      bool  `json:"use_available_plan"`
}

func Number3(arg *SardineCoverStrategy) {

	last = arg

}
