package colorized

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type colorCode = int

const (
	Black colorCode = 30
	Red   colorCode = 31
	Blue  colorCode = 34
)

func Color(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		after, err := colorize(text)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(w, "%s\n", after)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func colorize(str string) (after string, err error) {
	target := strings.TrimSpace(str)
	if strings.HasPrefix(target, "--- FAIL:") || strings.HasPrefix(target, "FAIL\t") {
		return colterm(str, Red), nil
	}

	// do nothing
	return str, nil
}

func colterm(str string, colorCode colorCode) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, str)
}
