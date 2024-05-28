package randstr

import (
	"math/rand"
)

type Randstr struct {
	length int
	rs     []rune
}

func New(options ...Option) *Randstr {
	r := &Randstr{
		length: 0,
	}

	if options == nil {
		options = append(options, WithDigit(), WithLetter())
	}

	for _, opt := range options {
		opt(r)
	}
	return r
}

func (r *Randstr) Random(length int) string {
	if length <= 0 {
		length = 1
	}

	r.length = length
	rs := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		r := r.rs[rand.Intn(len(r.rs))]
		rs = append(rs, r)
	}
	return string(rs)
}
