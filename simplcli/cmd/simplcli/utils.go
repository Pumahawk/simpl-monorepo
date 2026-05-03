package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"sync"
)

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

type SearchMultiProjectAsyncCmd[S any, T any] struct {
	Name        string
	Search      S
	FlagsSetter func(*flag.FlagSet)
	Validator   func() error
	RowConsumer func(string, *T)
	SortFunc    func([]T, int, int) bool
	ApiFunc     func(self *SearchMultiProjectAsyncCmd[S, T], projectId string, search *S) ([]T, error)
}

func (s *SearchMultiProjectAsyncCmd[S, T]) CName() string {
	return s.Name
}
func (s *SearchMultiProjectAsyncCmd[S, T]) CRun(args []string) (any, error) {
	var search *S
	searchT := reflect.TypeOf(s.Search)
	reflect.ValueOf(&search).Elem().Set(reflect.New(searchT))

	fl := flag.NewFlagSet("", flag.ExitOnError)
	structFlag(fl, search)
	if s.FlagsSetter != nil {
		s.FlagsSetter(fl)
	}
	fl.Parse(args)
	projectIds := prIdsDemux.demux(fl.Args())

	if len(projectIds) == 0 {
		return nil, fmt.Errorf("missing project ids")
	}

	if s.Validator != nil {
		if err := s.Validator(); err != nil {
			return nil, err
		}
	}

	// Prepare response object container
	type resTy struct {
		name string
		rs   []T
		err  error
	}
	wg := sync.WaitGroup{}
	ress := make(chan resTy)

	// Retrieve all pipelines async
	for _, projectId := range projectIds {
		projectId = prIds.Get(projectId)
		name := projectId
		wg.Go(func() {
			rs, err := s.ApiFunc(s, projectId, search)
			ress <- resTy{
				name: name,
				rs:   rs,
				err:  err,
			}
		})
	}
	go func() {
		defer close(ress)
		wg.Wait()
	}()

	// Collect all response items
	// Rewrite pipeline name with project name.
	items := make([]T, 0, len(projectIds)*20)
	for r := range ress {
		if r.err != nil {
			fmt.Fprintf(os.Stderr, "projectId %s: %s\n", r.name, r.err)
		}
		if !reflect.ValueOf(r.rs).IsNil() {
			for i := range r.rs {
				if s.RowConsumer != nil {
					s.RowConsumer(r.name, &r.rs[i])
				}
				rv := reflect.ValueOf(r.rs[i])
				items = append(items, rv.Interface().(T))
			}
		}
	}

	// Sort items by updated_at
	if s.SortFunc != nil {
		sort.Slice(items, func(i, j int) bool {
			return s.SortFunc(items, i, j)
		})
	}

	return items, nil
}
