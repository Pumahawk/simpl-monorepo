package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"simplcli/internal/simpl"
	"text/tabwriter"
)

var ApiKeypairsSearchCmd = command{
	Name: "keypair:search",
	Func: func(s ...string) int8 {

		var page, size int

		// Set search parameters with flags
		search := &simpl.KeypairsSearch{}
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.StringVar(&search.Name, "name", "", "")
		fs.IntVar(&page, "page", 0, "")
		fs.IntVar(&size, "size", 10, "")
		fs.Parse(s)

		c := simpl.NewAuthorityClient()
		r, err := c.Keypairs(page, size, search)
		if err != nil {
			log.Fatalf("unable to make keypair: %s", err)
		}

		fmt.Printf("page: %d\n", page)
		fmt.Printf("size: %d\n", size)
		w := tabwriter.NewWriter(os.Stdout, 6, 2, 2, ' ', 0)
		fmt.Fprintf(w, "%s\t%s\n", "Id", "Name")
		for _, k := range r.Items {
			fmt.Fprintf(w, "%s\t%s\n", k.Id, k.Name)
		}
		w.Flush()
		return 0
	},
}
