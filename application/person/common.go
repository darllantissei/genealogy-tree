package person

import (
	"context"
	"errors"
)

func (p *PersonService) checkDependencyInjection(ctx context.Context) error {

	const (
		dependecyInjectionPendingPersistenceDB = "dependecy database pending for person"
	)

	if p.PersistenceDB == nil {
		return errors.New(dependecyInjectionPendingPersistenceDB)
	}

	return nil
}
