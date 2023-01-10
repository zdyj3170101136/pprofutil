package main

import (
	"debug/elf"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/parca-dev/parca/pkg/symbol/demangle"
	"github.com/parca-dev/parca/pkg/symbol/elfutils"
)

func main() {
	f, err := elf.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	debug, _ := elfutils.NewDebugInfoFile(f, demangle.NewDemangler("simple", false))
	for _, pc := range pcs {
		debug.SourceLines(pc)
	}
	log.Println(time.Now().Sub(t))
	pprof.StopCPUProfile()
	t = time.Now()
	for _, pc := range pcs {
		debug.SourceLines(pc)
	}
	log.Println(time.Now().Sub(t))
}
