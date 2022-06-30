package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Main(t *testing.T) {
	main()
}

func Test_getEntry(t *testing.T) {
	input := []string{
		"2-5",
		"z:",
		"zzztvz",
	}
	output := processData(input)
	fmt.Printf("%+v", output)
}

func Test_valid(t *testing.T) {
	tt := []struct {
		name     string
		input    data
		expected bool
	}{
		{
			name: "should pass",
			input: data{
				policy: policy{
					low:   2,
					high:  8,
					value: "d",
				},
				password: "pddzddkdvqgxndd",
			},
			expected: true,
		},
		{
			name: "should fail",
			input: data{
				policy: policy{
					low:   8,
					high:  8,
					value: "d",
				},
				password: "pddzddkdvqgxndd",
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(t.Name(), func(t *testing.T) {
			_, ok := valid(tc.input)
			assert.Equal(t, ok, tc.expected)
		})
	}

}
