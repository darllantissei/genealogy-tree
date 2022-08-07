package factoryperson

import (
	"sync"

	servicerelationship "github.com/darllantissei/genealogy-tree/application/relationship"
	dbconnectionrelationship "github.com/darllantissei/genealogy-tree/adapters/db/sqlite/relationship"
)

type FactoryRelationship struct {
	serviceRelationship *servicerelationship.RelationshipService
	dbconnection  *dbconnectionrelationship.RelationshipDB
	lockFlow      *sync.Mutex
}

func (fr *FactoryRelationship) NewDataBaseConnection() *dbconnectionrelationship.RelationshipDB {

	if fr.lockFlow == nil {

		fr.lockFlow = &sync.Mutex{}

		fr.lockFlow.Lock()

		defer func() {

			fr.lockFlow.Unlock()

			fr.lockFlow = nil
		}()

	}

	if fr.dbconnection == nil {
		fr.dbconnection = dbconnectionrelationship.NewRelationshipDB()
	}

	return fr.dbconnection
}

func (fr *FactoryRelationship) NewService() *servicerelationship.RelationshipService {

	if fr.lockFlow == nil {

		fr.lockFlow = &sync.Mutex{}

		fr.lockFlow.Lock()

		defer func() {

			fr.lockFlow.Unlock()

			fr.lockFlow = nil
		}()

	}

	if fr.serviceRelationship == nil {

		fr.serviceRelationship = &servicerelationship.RelationshipService{
			PersistenceDB: fr.NewDataBaseConnection(),
		}

	}

	return fr.serviceRelationship
}
