package person

import (
	"context"

	"github.com/darllantissei/genealogy-tree/application/model"
)

type IPersonService interface {
	Record(ctx context.Context, prsn model.Person) (model.Person, error)
	Fetch(ctx context.Context, prsn model.Person) (model.Person, error)
}

type IPersonWriteDB interface {
	Insert(ctx context.Context, prsn model.Person) (model.Person, error)
	Update(ctx context.Context, prsn model.Person) error
}

type IPersonReadDB interface {
	Get(ctx context.Context, prsn model.Person) (model.Person, error)
}

type IPersonPersistenceDB interface {
	IPersonWriteDB
	IPersonReadDB
}
