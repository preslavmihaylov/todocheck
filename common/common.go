// Package common contains a bunch of small utility functions used throughout different independent other packages
package common

import (
	"fmt"
	"strings"
)

// RemoveEmptyTokens of a splitted string. Tokens such as "" are removed.
func RemoveEmptyTokens(ss []string) []string {
	res := []string{}
	for _, s := range ss {
		if s != "" {
			res = append(res, s)
		}
	}

	return res
}

// ArrayAsRegexAnyMatchExpression converting array to regexp string for matching any of elements
func ArrayAsRegexAnyMatchExpression(todos []string) string {
	if len(todos) == 0 {
		panic("Empty list of todo strings")
	}
	return fmt.Sprintf("(?:%s)", strings.Join(todos, "|"))
}
