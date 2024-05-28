package randstr_test

import (
	"fmt"
	"github.com/codepkgs/gokit/estring/randstr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandstr(t *testing.T) {
	ranLen := 16
	r := randstr.New(randstr.WithDigit(), randstr.WithLetter())
	s := r.Random(ranLen)
	assert.Equal(t, ranLen, len(s))
	fmt.Println(s)
}

func ExampleRandstr() {
	r := randstr.New(randstr.WithDigit(), randstr.WithLetter())
	s := r.Random(10)
	fmt.Println(s)
}
