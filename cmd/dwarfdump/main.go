// dwarfdump print all dwarf entry.
//
// usage: dwarfdump filepath
package main

import (
	"debug/dwarf"
	"debug/elf"
	"log"
	"os"
)

func main() {
	f, err := elf.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	d, err := f.DWARF()
	if err != nil {
		log.Fatal(err)
	}
	r := d.Reader()
	var lastE = &dwarf.Entry{}
	for e, err := r.Next(); err == nil; e, err = r.Next() {
		if e != nil {
			log.Println(e)
		}
		// there is no more section
		if lastE == nil && e == nil {
			break
		}
		lastE = e
	}
}
