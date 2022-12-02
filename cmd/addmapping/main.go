// addmapping add mapping to pprof file.
//
// usage: addmapping pprof filepath
//
// runtime/pprof use /proc/self/maps to get address mapping.
// in mac, the mapping is fake:
//
//	Mappings
//	1: 0x0/0x0/0x0   [FN]
//
// we add the mapping to make the disasm work right.
package main

import (
	"log"
	"math"
	"os"

	"github.com/google/pprof/profile"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	p, err := profile.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	m := profile.Mapping{
		ID:           1,
		Start:        0,
		Limit:        math.MaxUint64,
		Offset:       0,
		File:         os.Args[2],
		BuildID:      "",   // TODO
		HasFunctions: true, // consistent with go [FN]
	}
	p.Mapping[0] = &m
	f, err = os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	p.Write(f)
	f.Sync()
	f.Close()
}
