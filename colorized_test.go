package colorized

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestColor(t *testing.T) {
	buf := new(bytes.Buffer)

	r, err := os.Open("testdata/reviewdog_go_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = Color(r, buf)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(buf.String(), "=== RUN") {
		t.Fatal("wrong")
	}

	fmt.Println(buf.String())
}
