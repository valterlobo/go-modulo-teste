package main

import (
	"fmt"

	"lobo.tech/anotacao/repository"
)

func main() {
	anotacaoRespository := repository.NewAnotacaoRepository()
	anotacao, err := anotacaoRespository.GetByID("12233")

	if err == nil {
		fmt.Println(anotacao.UUID, anotacao.Text)
	} else {
		fmt.Println(err)
	}

}
