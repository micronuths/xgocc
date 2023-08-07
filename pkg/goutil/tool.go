package goutil

import (
	"bufio"
	"io"
)

type ReadFunc func(line string) error

func ForEachLine(r io.Reader, function ReadFunc) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		function(scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		if err == io.EOF {
			return nil
		} else {
			return err
		}

	}
	return nil
}
