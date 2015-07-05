package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/JeremyOT/address/lookup"
)

var v4 = flag.Bool("4", false, "Prefer IPv4 addresses.")
var interfaceName = flag.String("i", "", "The network interface to use for address lookup.")
var remoteHost = flag.String("r", "", "The remote host to use for address lookup.")
var tcpAddr = flag.Bool("tcp-addr", false, "Find and print a currently unused TCP address.")
var udpAddr = flag.Bool("udp-addr", false, "Find and print a currently unused UDP address.")
var tcpPort = flag.Bool("tcp-port", false, "Find and print a currently unused TCP port.")
var udpPort = flag.Bool("udp-port", false, "Find and print a currently unused UDP port.")

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
	if *tcpAddr {
		addr, err := FindOpenTCPAddress(*interfaceName, *v4)
		if err != nil {
			panic(err)
		}
		fmt.Println(addr)
		return
	}
	if *udpAddr {
		addr, err := FindOpenUDPAddress(*interfaceName, *v4)
		if err != nil {
			panic(err)
		}
		fmt.Println(addr)
		return
	}
	if *tcpPort {
		port, err := FindOpenTCPPort(*interfaceName, *v4)
		if err != nil {
			panic(err)
		}
		fmt.Println(port)
		return
	}
	if *udpPort {
		port, err := FindOpenUDPPort(*interfaceName, *v4)
		if err != nil {
			panic(err)
		}
		fmt.Println(port)
		return
	}
	if *remoteHost != "" {
		if addr, err := LocalAddress(*remoteHost); err == nil {
			fmt.Println(addr)
			return
		} else {
			panic(err)
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
		panic(err)
	}
	fmt.Println(address)
}
