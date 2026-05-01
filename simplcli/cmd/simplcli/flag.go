package main

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

func structFlag(fl *flag.FlagSet, str any) {
	v := reflect.ValueOf(str)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	t := v.Type()
	for i := range t.NumField() {
		ft := t.Field(i)
		fv := v.Field(i)
		name := strings.ToLower(ft.Name)
		safeDefineFlag(fl, name, fv) // Ignore error: already defined flag, not supported flags
	}
}

func safeDefineFlag(fl *flag.FlagSet, name string, fv reflect.Value) error {
	if v := fl.Lookup(name); v != nil {
		return fmt.Errorf("already defined flag %q", name)
	}
	switch fv.Kind() {
	case reflect.String:
		fl.StringVar(fv.Addr().Interface().(*string), name, "", "")
	case reflect.Bool:
		fl.BoolVar(fv.Addr().Interface().(*bool), name, false, "")
	case reflect.Int:
		fl.IntVar(fv.Addr().Interface().(*int), name, 0, "")
	default:
		return fmt.Errorf("not supported kind %q field name %q", name, fv.Kind())
	}

	return nil
}
