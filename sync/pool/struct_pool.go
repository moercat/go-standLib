package main

import "sync"

type SNodeRedirectData struct {
	NodeKey   string
	Cover     string
	View      string
	Region    string
	SNodeName string
	Type      int
	Value     int64
}

type NotRedirectPoolType struct {
	*sync.Pool
}

func NewElt() (nPool NotRedirectPoolType) {
	nPool = NotRedirectPoolType{
		&sync.Pool{New: func() interface{} {
			return new(SNodeRedirectData)
		}},
	}
	return
}

func (s NotRedirectPoolType) GetElt() *SNodeRedirectData {
	return s.Get().(*SNodeRedirectData)
}

func (s NotRedirectPoolType) PutElt(data *SNodeRedirectData) {

	data = new(SNodeRedirectData)

	s.Put(data)
	return
}

func (s NotRedirectPoolType) PutEltNoNew(data *SNodeRedirectData) {

	data.NodeKey = ""
	data.Cover = ""
	data.View = ""
	data.Region = ""
	data.SNodeName = ""
	data.Type = 0
	data.Value = 0

	s.Put(data)
	return
}

func (s NotRedirectPoolType) PutNil(data *SNodeRedirectData) {

	return
}

var NotRedirectPool = NewElt()
