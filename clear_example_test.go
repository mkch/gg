package gg_test

import (
	"sync"

	"github.com/mkch/gg"
)

type Request struct {
	Field1      int
	Field2      string
	SecretToken string
}

var requestPool = sync.Pool{New: func() any {
	return &Request{}
}}

func NewRequest(f1 int, f2 string) *Request {
	req := requestPool.Get().(*Request)
	// Clear all fields to avoid leaking SecretToken
	// It is safer than manually resetting each field
	// as the struct may evolve over time.
	gg.Clear(req)
	req.Field1 = f1
	req.Field2 = f2
	return req
}
