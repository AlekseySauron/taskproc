package testFib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "one",
			n:    1,
			want: 1,
		},
		{
			name: "two",
			n:    2,
			want: 3,
		},
	}

	for _, tc := range testCases {

		// v := Fib(tc.n)
		// if v != tc.want {
		// 	t.Error(tc.name)
		// }
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Fib(tc.n))
		})
	}

	// var curDir string

	// curDir, _ = os.Getwd()

	// fmt.Println("curDir =", curDir)

	// err, db =  initBD()
	// if err {
	//return
	// }
}

func Fib(n int) int {

	return n * n

}

func TestMy(t *testing.T) {

	// var curDir string

	// curDir, _ = os.Getwd()

	//fmt.Println("curDir =", curDir)
	//fmt.Printf("t: %v\n", curDir)

}
