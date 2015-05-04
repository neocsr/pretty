package format_test

import (
	"fmt"
	"testing"

	"github.com/neocsr/pretty/format"
)

func TestPrettify(t *testing.T) {
	tokens, err := format.Prettify("http://test.com:8080/search?q=example")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tokens)
}
