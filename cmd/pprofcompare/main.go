// package pprofcompare compare two profile's all location's address
//
// usage: pprofcompare basefile comparefile
package main

import (
	"log"
	"os"

	"github.com/google/pprof/profile"
)

func main() {
	base, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	cmp, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	baseP, err := profile.Parse(base)
	if err != nil {
		log.Fatal(err)
	}
	cmpP, err := profile.Parse(cmp)
	if err != nil {
		log.Fatal(err)
	}

	for i := range cmpP.Location {
		address := cmpP.Location[i].Address
		var ok bool
		for j := range baseP.Location {
			if baseP.Location[j].Address == address {
				ok = true
				break
			}
		}

		if !ok {
			log.Printf("compare profile have an address not in base profile: %x\n", address)
		}
	}

	for i := range baseP.Location {
		address := baseP.Location[i].Address
		var ok bool
		for j := range cmpP.Location {
			if cmpP.Location[j].Address == address {
				ok = true
				break
			}
		}

		if !ok {
			log.Printf("base profile have an address not in compare profile: %x\n", address)
		}
	}
}
