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

		// Set search parameters with flags
		search := &simpl.KeypairsSearch{}
		fs := flag.NewFlagSet("", flag.ExitOnError)
		fs.StringVar(&search.Name, "name", "", "")
		fs.Parse(s)

		c := simpl.NewAuthorityClient()
		r, err := c.Keypairs(0, 10, search)
		if err != nil {
			log.Fatalf("unable to make keypair: %s", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 6, 2, 2, ' ', 0)
		fmt.Fprintf(w, "%s\t%s\n", "Id", "Name")
		for _, k := range r.Items {
			fmt.Fprintf(w, "%s\t%s\n", k.Id, k.Name)
		}
		w.Flush()
		return 0
	},
}
