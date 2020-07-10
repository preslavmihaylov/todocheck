package lines

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type lineCallback func(filename, line string, linecnt int) error

// TraversePath and perform a callback on each line in each file
func TraversePath(path string, ignoredPaths, supportedFileExtensions []string, callback lineCallback) error {
	return filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("couldn't traverse %s: %w", file, err)
		}

		if info.IsDir() && isIgnored(ignoredPaths, file) {
			fmt.Println("Skipping ignored dir", file)
			return filepath.SkipDir
		} else if info.IsDir() || !isSupported(supportedFileExtensions, file) {
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

func isIgnored(ignoredPaths []string, path string) bool {
	if isHidden(path) {
		return true
	}

	for _, ignoredPath := range ignoredPaths {
		isMatch, err := filepath.Match(ignoredPath, path)
		if err != nil {
			log.Fatalf("Couldn't process glob pattern %s for path %s: %s", ignoredPath, path, err)
		}

		if isMatch {
			return true
		}
	}

	return false
}

func isSupported(supportedExtensions []string, file string) bool {
	for _, ext := range supportedExtensions {
		if filepath.Ext(file) == ext {
			return true
		}
	}

	return false
}

func isHidden(path string) bool {
	return len(path) > 1 && path[:2] != "./" && path[:2] != ".." && path[0] == byte('.')
}
