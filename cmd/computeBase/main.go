package main

import (
	"log"
	"os"
	"strconv"

	"github.com/google/pprof/profile"

	"github.com/zdyj3170101136/pprofutil/cmd/computeBase/objectfile"
)

func main() {
	name := os.Args[1]
	start, limit, offset := parseOx(os.Args[2]), parseOx(os.Args[3]), parseOx(os.Args[4])
	addr := parseOx(os.Args[5])
	o, err := objectfile.Open(name, &profile.Mapping{
		Start:  start,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		log.Fatal(err)
	}
	result, err := o.ObjAddr(addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result: %x \n", result)
	log.Printf("base: %x \n", addr-result)
}

func parseOx(s string) uint64 {
	res, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
