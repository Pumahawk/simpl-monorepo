package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"simplcli/internal/simpl"
	"slices"
	"strings"
	"text/tabwriter"
)

type SearchFunc[T any, R any] func(page, size int, filter T) (R, error)

var ApiKeypairsSearchCmd = command{
	Name: "keypair:search",
	Func: func(s ...string) int8 {
		// Set search parameters with flags
		search := &simpl.KeypairsSearch{}
		c := simpl.NewAuthorityClient()
		return searchBaseCmd(search, c.Keypairs, s)
	},
}

func searchBaseCmd[T any, R any](search T, searchFunc SearchFunc[T, R], s []string) int8 {
	// Set search parameters with flags
	fs := flag.NewFlagSet("", flag.ExitOnError)
	page, size := searchFlag(fs, "sh", search)
	fieldsFlag := fs.String("fields", "", "")
	fs.Parse(s)

	fieldsS := *fieldsFlag
	var fieldsArgs []string
	if fieldsS != "" {
		fieldsArgs = strings.Split(*fieldsFlag, ",")
	}

	r, err := searchFunc(*page, *size, search)
	if err != nil {
		log.Fatalf("unable to make search: %s", err)
	}

	// Get printable fields, default print all fields
	rf := reflect.ValueOf(r).Elem().FieldByName("Items")
	rt := rf.Type().Elem()
	structFields := getFieldsName(rt)
	fields := make([]string, 0, len(structFields))
	if len(fieldsArgs) == 0 {
		fields = structFields
	} else {
		for _, fn := range fieldsArgs {
			if slices.Contains(structFields, fn) {
				fields = append(fields, fn)
			}
		}
	}

	// Print table view of fields
	fmt.Printf("page: %d\n", *page)
	fmt.Printf("size: %d\n", *size)
	w := tabwriter.NewWriter(os.Stdout, 6, 2, 2, ' ', 0)
	w.Write([]byte(strings.Join(fields, "\t") + "\n"))
	for i := 0; i < rf.Len(); i++ {
		item := rf.Index(i)
		var itemsS []string
		for _, fieldName := range fields {
			fmts := "%v"
			itemField := item.FieldByName(fieldName)
			itemValue := itemField.Interface()
			if itemField.Kind() == reflect.String {
				fmts = "%q"
			}
			itemsS = append(itemsS, fmt.Sprintf(fmts, itemValue))
		}
		w.Write([]byte(strings.Join(itemsS, "\t") + "\n"))
	}
	w.Flush()

	return 0
}

func searchFlag(fl *flag.FlagSet, prefix string, search any) (*int, *int) {
	rf := reflect.ValueOf(search)
	if rf.Kind() == reflect.Pointer {
		rf = rf.Elem()
	}
	rt := rf.Type()
	for i := range rf.NumField() {
		f := rf.Field(i)
		name := rt.Field(i).Tag.Get("search")
		if name == "" {
			name = rt.Field(i).Name
		}
		flagName := prefix + "-" + name
		if prefix == "" {
			flagName = name
		}
		fl.StringVar(f.Addr().Interface().(*string), flagName, "", "")
	}
	page := fl.Int("page", 0, "")
	size := fl.Int("size", 10, "")
	return page, size
}

func getFieldsName(t reflect.Type) []string {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("only supported valid struct")
	}
	f := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f = append(f, t.Field(i).Name)
	}
	return f
}
