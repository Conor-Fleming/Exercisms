package main

import (
	"fmt"
	"ledger/ledger"
	"log"
)

var entries = []ledger.Entry{
	{
		Date:        "2015-01-01",
		Description: "Buy present",
		Change:      -1000,
	},
}

func main() {
	result, err := ledger.FormatLedger("USD", "en-US", entries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
