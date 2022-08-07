package relationship

import (
	"context"
	"errors"
)

func (p *RelationshipService) checkDependencyInjection(ctx context.Context) error {

	const (
		dependecyInjectionPendingPersistenceDB = "dependecy database pending for relationship"
		dependecyInjectionPendingPersonService = "dependecy of person service is pending"
	)

	if p.PersistenceDB == nil {
		return errors.New(dependecyInjectionPendingPersistenceDB)
	}

	if p.PersonService == nil {
		return errors.New(dependecyInjectionPendingPersonService)
	}

	return nil
}
