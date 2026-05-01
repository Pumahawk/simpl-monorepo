package main

import "strings"

type projectNameSvT map[string]string

func (p projectNameSvT) Get(key string) string {
	v := p[key]
	if v == "" {
		return key
	} else {
		return v
	}
}

func getFields(fieldsS string) (fields []string) {
	if fieldsS != "" {
		fields = strings.Split(fieldsS, ",")
	}
	return
}
