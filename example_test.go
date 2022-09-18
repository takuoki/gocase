package gocase_test

import (
	"fmt"

	"github.com/takuoki/gocase"
)

func Example() {

	strs := []string{"IpAddress", "defaultDnsServer", "JsonFile", "CsvFile"}

	fmt.Println("-- Default converter --")
	for _, str := range strs {
		fmt.Printf("%s --To-> %s --Revert-> %s\n", str, gocase.To(str), gocase.Revert(gocase.To(str)))
	}
	fmt.Println()

	fmt.Println("-- Custom converter --")
	conv1, _ := gocase.New(gocase.WithInitialisms("JSON", "CSV"))
	for _, str := range strs {
		fmt.Printf("%s --To-> %s --Revert-> %s\n", str, conv1.To(str), conv1.Revert(conv1.To(str)))
	}
	fmt.Println()

	fmt.Println("-- Custom converter (add to default) --")
	initialisms := append([]string{"JSON", "CSV"}, gocase.DefaultInitialisms...)
	conv2, _ := gocase.New(gocase.WithInitialisms(initialisms...))
	for _, str := range strs {
		fmt.Printf("%s --To-> %s --Revert-> %s\n", str, conv2.To(str), conv2.Revert(conv2.To(str)))
	}

	// Output:
	// -- Default converter --
	// IpAddress --To-> IPAddress --Revert-> IpAddress
	// defaultDnsServer --To-> defaultDNSServer --Revert-> defaultDnsServer
	// JsonFile --To-> JSONFile --Revert-> JsonFile
	// CsvFile --To-> CsvFile --Revert-> CsvFile
	//
	// -- Custom converter --
	// IpAddress --To-> IpAddress --Revert-> IpAddress
	// defaultDnsServer --To-> defaultDnsServer --Revert-> defaultDnsServer
	// JsonFile --To-> JSONFile --Revert-> JsonFile
	// CsvFile --To-> CSVFile --Revert-> CsvFile
	//
	// -- Custom converter (add to default) --
	// IpAddress --To-> IPAddress --Revert-> IpAddress
	// defaultDnsServer --To-> defaultDNSServer --Revert-> defaultDnsServer
	// JsonFile --To-> JSONFile --Revert-> JsonFile
	// CsvFile --To-> CSVFile --Revert-> CsvFile
}
