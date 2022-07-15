package main

import (
	"fmt"
)

func main() {
	rep := NewRepository("tx")
	entity, _ := rep.GetByID(4)
	tx := entity.(*Transaction)
	//tx := entity.GetTX()
	fmt.Println(tx.Amount)

}

type IEntity interface {
}

type Transaction struct {
	ID     int
	TxHash string
	Amount uint64
}

type Repository interface {
	GetByID(ID int) (IEntity, error)
}

func (t *Transaction) GetTX() *Transaction {
	return t
}

func NewRepository(repositoryType string) Repository {
	switch repositoryType {
	case "tx":
		return NewTransactionRepository()
	}
	return nil
}

type TransactionRepository struct {
}

func NewTransactionRepository() Repository {
	return &TransactionRepository{}
}

func (t *TransactionRepository) GetByID(ID int) (IEntity, error) {
	return &Transaction{TxHash: "gddfd", Amount: 5}, nil
}
