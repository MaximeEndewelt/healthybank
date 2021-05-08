package main

import (
	"fmt"
	"healthybank/clients"
)

func main() {
	println("Hello")
	client := clients.New()
	if err := client.GetAccountInformation(); err != nil {
		fmt.Errorf("%s", err)
	}
}
