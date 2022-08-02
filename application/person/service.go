package person

import (
	"context"
	"strings"

	"github.com/darllantissei/genealogy-tree/application/common"
	"github.com/darllantissei/genealogy-tree/application/model"
)

type PersonService struct {
	PersistenceDB IPersonPersistenceDB
}

func (p *PersonService) Record(ctx context.Context, prsn model.Person) (model.Person, error) {

	var (
		err error
	)

	err = p.checkDependencyInjection(ctx)

	if err != nil {
		return model.Person{}, common.BuildError(ctx, []string{err.Error()})
	}

	msgErr := p.checkDataPerson(ctx, prsn)

	if err != nil {
		return model.Person{}, common.BuildError(ctx, msgErr)
	}

	if strings.EqualFold(prsn.ID, "") {
		prsn, err = p.PersistenceDB.Insert(ctx, prsn)
	} else {
		err = p.PersistenceDB.Update(ctx, prsn)
	}

	return prsn, err
}
func (p *PersonService) Fetch(ctx context.Context, prsn model.Person) (model.Person, error) {

	var (
		err error
	)

	err = p.checkDependencyInjection(ctx)

	if err != nil {
		return model.Person{}, common.BuildError(ctx, []string{err.Error()})
	}

	prsn, err = p.PersistenceDB.Get(ctx, prsn)

	if err != nil {
		return model.Person{}, common.BuildError(ctx, []string{err.Error()})
	}

	return prsn, nil
}
