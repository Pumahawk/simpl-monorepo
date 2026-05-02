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

type projectIdsDemux map[string][]string

func (p projectIdsDemux) demux(v []string) []string {
	strs := make([]string, 0, len(v))
	for _, s := range v {
		if prs, ok := p[s]; ok {
			strs = append(strs, prs...)
		} else {
			strs = append(strs, s)
		}
	}
	return strs
}
