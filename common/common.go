// Package common contains a bunch of small utility functions used throughout different independent other packages
package common

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
