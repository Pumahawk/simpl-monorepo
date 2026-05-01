package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

type TableView struct {
	w *tabwriter.Writer
}

func StdTableWriter() TableView {
	w := os.Stdout
	tw := tabwriter.NewWriter(w, 5, 2, 2, ' ', 0)
	return TableView{tw}
}

type RenderOpt struct {
	Fields []string
}

func (t *TableView) RenderList(opt *RenderOpt, model any) error {
	// Validation model, retrieve slice
	rv := reflect.ValueOf(model)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		t := rv.Type()
		if _, ok := t.FieldByName("Items"); !ok {
			return fmt.Errorf("struct doesn't contains Items field")
		}
		rv = rv.FieldByName("Items")
	}
	if rv.Kind() != reflect.Slice {
		return fmt.Errorf("not supported kind %s", rv.Kind())
	}

	rt := rv.Type().Elem()

	// Define fields
	var fields []string
	if opt != nil {
		fields = opt.Fields
	}
	if len(fields) == 0 {
		fields = make([]string, 0, rt.Size())
		for i := range rt.NumField() {
			fields = append(fields, rt.Field(i).Name)
		}
	}

	// Write table header
	t.w.Write([]byte(strings.Join(fields, "\t") + "\n"))

	// Write table body
	for i := range rv.Len() {
		rv := rv.Index(i)
		row := make([]string, 0, rt.NumField())
		for _, field := range fields {
			if _, ok := rt.FieldByName(field); ok {
				value := rv.FieldByName(field)
				fmts := "%v"
				if value.Kind() == reflect.String {
					fmts = "%q"
				}
				row = append(row, fmt.Sprintf(fmts, value))
			}
		}
		t.w.Write([]byte(strings.Join(row, "\t") + "\n"))
	}

	return t.w.Flush()
}

func (t *TableView) RenderValue(opt *RenderOpt, model any) error {
	// Validation model
	rv := reflect.ValueOf(model)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("unsupported type %s", rv.Kind())
	}

	// Retrieve header
	rt := rv.Type()
	var fields []string
	if opt != nil {
		fields = opt.Fields
	}
	if len(fields) == 0 {
		for i := range rt.NumField() {
			fields = append(fields, rt.Field(i).Name)
		}
	}

	// Retrieve and print body
	for _, f := range fields {
		if ft, ok := rt.FieldByName(f); ok {
			if !rv.FieldByName(f).IsZero() {
				name := ft.Name
				value := rv.FieldByName(f)
				fmts := "%s\t%v\n"
				if value.Kind() == reflect.String {
					fmts = "%s\t%q\n"
				}
				fmt.Fprintf(t.w, fmts, name, value.Interface())
			}
		}
	}
	t.w.Flush()
	return nil
}
