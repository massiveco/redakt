package main

import "strings"

type redactFlags []string

func (rf *redactFlags) String() string {
	return strings.Join(*rf, ", ")
}

func (rf *redactFlags) Set(value string) error {
	*rf = append(*rf, value)
	return nil
}

type manifestStub struct {
	Kind     string
	Metadata manifestMetadata
}

type manifestMetadata struct {
	Name string
}
