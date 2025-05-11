package templates

import (
	"sitegenerator/infra/testdata"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	cases := []struct {
		expected  string
		inputs    []string
		separator string
	}{
		// Multiple inputs, space separator
		{
			expected:  "The Old New Thing",
			inputs:    []string{"The", "Old", "New", "Thing"},
			separator: " ",
		},
		// Multiple inputs, comma separator
		{
			expected:  "Hello,World",
			inputs:    []string{"Hello", "World"},
			separator: ",",
		},
		// Single input
		{
			expected:  "Hello",
			inputs:    []string{"Hello"},
			separator: ",",
		},
		// Empty inputs
		{
			expected:  "",
			inputs:    []string{},
			separator: ",",
		},
	}

	for _, c := range cases {
		name := "TestJoin_" + c.expected
		t.Run(name, func(t *testing.T) {
			actual := join(c.inputs, c.separator)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestAssetHash(t *testing.T) {
	cases := []struct {
		expected string
		urlPath  string
	}{
		{
			expected: "/images/1.gif?v=c571d09b80b8b73e43fd2fec1951ea2e",
			urlPath:  "/images/1.gif",
		},
		{
			expected: "/images/2.jpg?v=5e111f61ce6e24a5a921430285758d21",
			urlPath:  "/images/2.jpg",
		},
	}

	callbacks := CreateFuncCallbacks(testdata.ContentDir())
	for _, c := range cases {
		name := "TestAssetHash" + strings.ReplaceAll(c.urlPath, "/", "_")
		t.Run(name, func(t *testing.T) {
			actual := addAssetHash(callbacks, c.urlPath)
			assert.Equal(t, c.expected, actual)
		})
	}
}
