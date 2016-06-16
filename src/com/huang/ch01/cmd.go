package ch01

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        [] string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage() = printUsage

}

func printUsage() {
	fmt.Printf("Usage:%s [-options] class [args...]\n", os.Args[0])

}
