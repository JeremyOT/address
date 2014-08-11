package main

import (
	"flag"
	"fmt"
	. "github.com/JeremyOT/address/lookup"
	"os"
)

var v4 = flag.Bool("4", false, "Prefer IPv4 addresses.")
var interfaceName = flag.String("i", "", "The network interface to use for address lookup.")
var remoteHost = flag.String("r", "", "The remote host to use for address lookup.")

func main() {
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  Prints the first broadcast network address found for the current host.")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()
	if *remoteHost != "" {
		if addr, err := LocalAddress(*remoteHost); err == nil {
			fmt.Println(addr)
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	var address string
	var err error
	if *interfaceName == "" {
		address, err = GetAddress(*v4)
	} else {
		address, err = GetInterfaceAddress(*interfaceName, *v4)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(address)
}
