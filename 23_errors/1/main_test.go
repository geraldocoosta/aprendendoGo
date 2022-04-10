package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoma(t *testing.T) {
	assert := assert.New(t)

	soma := Soma(1, 2)
	assert.Equal(soma, 3)

}

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{[]int{10, 20}, 30},
		{[]int{1, 6}, 7},
		{[]int{1, 2}, 3},
	}

	for _, test := range tests {
		if result := Soma(test.data[0], test.data[1]); result != test.answer {
			t.Errorf("soma(%v) = %v, want %v", test.data, result, test.answer)
		}
	}
}

func ExampleSoma() {
	fmt.Println(Soma(10, 20))
	fmt.Println(Soma(10, 21))
	fmt.Println(Soma(10, 22))
	// Output:
	// 30
	// 31
	// 32
}
