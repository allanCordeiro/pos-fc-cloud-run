package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/infra/service/retrievecep/impl"
	"github.com/allanCordeiro/pos-fc-cloud-run/internal/usecase/cep"
)

func main() {
	searchCep := impl.NewViaCep(http.DefaultClient)
	usecase := cep.NewRetrieveUseCase(searchCep)
	cep := cep.Input{Zipcode: "04266-060"}
	output, err := usecase.Execute(context.TODO(), cep)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
