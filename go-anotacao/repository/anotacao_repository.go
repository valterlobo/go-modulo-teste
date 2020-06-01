package repository

import "lobo.tech/anotacao/model"

type AnotacaoRepository struct {
}

func NewAnotacaoRepository() *AnotacaoRepository {

	return &AnotacaoRepository{}

}

func (repository *AnotacaoRepository) GetByID(id string) (model.Anotacao, error) {

	var anotacao model.Anotacao = model.Anotacao{id, "hello texto"}

	return anotacao, nil

}
