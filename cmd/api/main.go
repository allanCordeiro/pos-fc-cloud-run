package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/infra/service/retrievecep/impl"
)

func main() {
	searchCep := impl.NewViaCep(http.DefaultClient)
	zipcode, err := searchCep.Retrieve(context.TODO(), "04266060")
	if err != nil {
		panic(err)
	}
	fmt.Println(zipcode)
}
