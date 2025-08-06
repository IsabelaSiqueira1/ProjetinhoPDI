package leituraescrita

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Collection interface {
	Insert(int)
	ForEach(func(int))
}

func Read(reader io.Reader, collection Collection) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		v, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("parse %q: %w", line, err)
		}
		collection.Insert(v)
	}
	return scanner.Err()
}

func Write(writer io.Writer, collection Collection) error {
	var err error
	collection.ForEach(func(v int) {
		if _, e := fmt.Fprintf(writer, "%d\n", v); e != nil {
			err = e
		}
	})
	return err
}
