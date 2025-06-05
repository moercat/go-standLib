package main

import "testing"

func BenchmarkGet1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := NotRedirectPool.GetElt()

		data = &SNodeRedirectData{
			NodeKey:   "1",
			Cover:     "1",
			View:      "1",
			Region:    "1",
			SNodeName: "1",
			Type:      1,
			Value:     1,
		}

		NotRedirectPool.PutEltNoNew(data)
	}
}

func BenchmarkGet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := NotRedirectPool.GetElt()

		data = &SNodeRedirectData{
			NodeKey:   "1",
			Cover:     "1",
			View:      "1",
			Region:    "1",
			SNodeName: "1",
			Type:      1,
			Value:     1,
		}

		NotRedirectPool.PutElt(data)
	}
}
