package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	assert := assert.New(t)

	soma := Soma(1, 2)
	assert.Equal(soma, 3)

}
