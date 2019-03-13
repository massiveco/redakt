package main

import "strings"

var sep = []byte("---")

func splitOnDashes(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.Index(string(data), "---"); i >= 0 {
		return i + 3, append(data[0:i], sep...), nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
