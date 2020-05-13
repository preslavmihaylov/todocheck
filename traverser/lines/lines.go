package lines

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type lineCallback func(filename, line string, linecnt int) error

// TraversePath and perform a callback on each line in each file
func TraversePath(path string, callback lineCallback) error {
	return filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if info.IsDir() || isIgnored(file) || filepath.Ext(file) != ".go" {
			return nil
		}

		err = traverseFile(file, callback)
		if err != nil {
			return fmt.Errorf("failed traversing file %s: %w", file, err)
		}

		return nil
	})
}

func traverseFile(filename string, callback lineCallback) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}

	var line string
	var linesRead int
	linecnt := 0

	reader := bufio.NewReader(bytes.NewReader(buf))
	for {
		linecnt++
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		linecnt += linesRead
		err = callback(filename, line, linecnt)
		if err != nil {
			return err
		}
	}

	if err != io.EOF {
		return fmt.Errorf("encountered error while traversing file %s: %w", filename, err)
	}

	return nil
}

func isIgnored(path string) bool {
	return path[0] == byte('.')
}
