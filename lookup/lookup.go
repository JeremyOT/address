package lookup

import (
	"fmt"
	"net"
	"strings"
)

// Make a connection to testUrl to infer the local network address. Returns
// the address of the interface that was used for the connection.
func LocalAddress(testUrl string) (host string, err error) {
	if strings.HasPrefix(testUrl, "http://") {
		testUrl = testUrl[7:]
	} else if strings.HasPrefix(testUrl, "https://") {
		testUrl = testUrl[8:]
	}
	if !strings.ContainsRune(testUrl, ':') {
		testUrl = testUrl + ":80"
	}
	if conn, err := net.Dial("udp", fmt.Sprintf("%s", testUrl)); err != nil {
		return "", err
	} else {
		defer conn.Close()
		host = strings.Split(conn.LocalAddr().String(), ":")[0]
	}
	return
}

// List all IP addresses found for the named interface
func InterfaceIPs(interfaceName string) (ips []net.IP, err error) {
	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return
	}
	addrs, err := iface.Addrs()
	if err != nil {
		return
	}
	ips = make([]net.IP, 0, len(addrs))
	for _, addr := range addrs {
		if ip, _, err := net.ParseCIDR(addr.String()); err == nil {
			ips = append(ips, ip)
		}
	}
	return
}

// List the first IP address found for the named interface
func InterfaceIP(interfaceName string) (ip net.IP, err error) {
	ips, err := InterfaceIPs(interfaceName)
	if err != nil {
		return
	}
	return ips[0], nil
}

// List the first IPv4 address found for the named interface
func InterfaceIPv4(interfaceName string) (ip net.IP, err error) {
	ips, err := InterfaceIPs(interfaceName)
	if err != nil {
		return
	}
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip, nil
		}
	}
	return
}

// List all interfaces matching the given flags. Or all interfaces if no flags
// are passed. E.g. FilterInterfaces(net.FlagUp, net.FlagBroadcast)
func FilterInterfaces(f ...net.Flags) (interfaces []net.Interface, err error) {
	allInterfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	interfaces = make([]net.Interface, 0, len(allInterfaces))
	for _, iface := range allInterfaces {
		include := true
		for _, flag := range f {
			if iface.Flags&flag == 0 {
				include = false
			}
		}
		if include {
			interfaces = append(interfaces, iface)
		}
	}
	return
}

// Get the first address found for the named interface, optionally returning
// only IPv4 addresses.
func GetInterfaceAddress(name string, filterIPv4 bool) (address string, err error) {
	var ip net.IP
	if filterIPv4 {
		if ip, err = InterfaceIPv4(name); err != nil {
			return
		}
	} else {
		if ip, err = InterfaceIP(name); err != nil {
			return
		}
	}
	return ip.String(), nil
}

// Get the first address found for the first broadcast interface, optionally returning
// only IPv4 addresses.
func GetAddress(filterIPv4 bool) (address string, err error) {
	interfaces, err := FilterInterfaces(net.FlagUp, net.FlagBroadcast)
	if err != nil {
		return
	}
	address, err = GetInterfaceAddress(interfaces[0].Name, filterIPv4)
	return
}
