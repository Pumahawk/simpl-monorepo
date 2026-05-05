package vw

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"
)

type TableView struct {
	w io.Writer
}

func NewTableWriter(w io.Writer) *TableView {
	return &TableView{w}
}

func (t *TableView) newTavWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(t.w, 5, 2, 2, ' ', 0)
}

func (t *TableView) Render(opt *RenderOpt, model any) error {
	// Validation model, retrieve slice
	rv := reflect.ValueOf(model)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Struct:
		return t.renderStruct(opt, model)
	case reflect.Slice:
		return t.renderList(opt, model)
	}
	return fmt.Errorf("unsuppported model")
}

func (t *TableView) renderList(opt *RenderOpt, model any) error {
	tw := t.newTavWriter()
	// Validation model, retrieve slice
	rv := reflect.ValueOf(model)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
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
	tw.Write([]byte(strings.Join(fields, "\t") + "\n"))

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
		tw.Write([]byte(strings.Join(row, "\t") + "\n"))
	}

	return tw.Flush()
}

func (t *TableView) renderStruct(opt *RenderOpt, model any) error {
	tw := t.newTavWriter()

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
				if value.Kind() == reflect.Pointer {
					value = value.Elem()
				}
				fmts := "%s\t%v\n"
				switch value.Kind() {
				case reflect.String:
					fmts = "%s\t%q\n"
					fmt.Fprintf(tw, fmts, name, value.Interface())
				case reflect.Slice:
					fmt.Fprintf(t.w, "[[%s]]\n", name)
					t.renderList(opt, value.Interface())
				case reflect.Struct:
					fmt.Fprintf(t.w, "[%s]\n", name)
					t.renderStruct(opt, value.Interface())
				default:
					fmt.Fprintf(tw, fmts, name, value.Interface())
				}
			}
		}
	}
	tw.Flush()
	return nil
}
