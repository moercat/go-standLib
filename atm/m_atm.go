package atm

import "sync/atomic"

func AAA() {
	notRedirectLimit := make(map[string]*int64)
	var a int64
	notRedirectLimit["AAA"] = &a

	if atomic.LoadInt64(notRedirectLimit["AAA"]) > 0 {
		return
	}

	if _, ok := notRedirectLimit["BBB"]; ok {
		atomic.SwapInt64(notRedirectLimit["BBB"], 1)
	}
	atomic.SwapInt64(notRedirectLimit["AAA"], 1)
	if atomic.LoadInt64(notRedirectLimit["AAA"]) > 0 {
		return
	}
}
