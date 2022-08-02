package relationship

import (
	"context"
	"errors"
)

func (p *RelationshipService) checkDependencyInjection(ctx context.Context) error {

	const (
		dependecyInjectionPendingPersistenceDB = "dependecy database pending for relationship"
	)

	if p.PersistenceDB == nil {
		return errors.New(dependecyInjectionPendingPersistenceDB)
	}

	return nil
}
