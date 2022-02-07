package models

import (
	"strings"
)

type StringList []string

func (l *StringList) Set(s string) error {
	*l = append(*l, s)
	return nil
}

func (l StringList) String() string {
	return strings.Join(l, ",")
}
