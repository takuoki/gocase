package gocase_test

import (
	"fmt"

	"github.com/takuoki/gocase"
)

func Example() {
	fmt.Println(gocase.To("IpAddress"))
	fmt.Println(gocase.Revert("defaultDNSServer"))
	
	// Output:
	// IPAddress
	// defaultDnsServer
}