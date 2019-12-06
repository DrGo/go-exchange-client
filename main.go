package main

import (
	"fmt"

	"github.com/fxpgr/go-exchange-client/api/private"
)

func main() {
	pcli, err := private.NewClient(private.PROJECT, "binance", func() (string, error) {
		return "", nil
	}, func() (s string, err error) {
		return "", nil
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(pcli.Balances())
}
