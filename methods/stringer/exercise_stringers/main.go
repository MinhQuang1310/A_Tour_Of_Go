package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// Make the IPAddr type implement fmt.Stringer
// to print the address as a dotted quad.
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

func (ip IPAddr) String() string {
	path := []string{
		fmt.Sprint(ip[0]),
		fmt.Sprint(ip[1]),
		fmt.Sprint(ip[2]),
		fmt.Sprint(ip[3]),
	}
	return strings.Join(path, ".")
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {

		fmt.Printf("%v: %v\n", name, ip)
	}
}
