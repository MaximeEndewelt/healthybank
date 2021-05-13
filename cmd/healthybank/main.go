package main

import (
	"healthybank.com/healthybank/api"
	"healthybank.com/healthybank/financial_project"
)

func main() {
	// client := clients.New()
	// if err := client.GetAccountInformation(); err != nil {
	// 	fmt.Errorf("%s", err)
	// }
	svc := financial_project.New()
	fpHandler := api.New(svc)

	fpHandler.Start()
}
