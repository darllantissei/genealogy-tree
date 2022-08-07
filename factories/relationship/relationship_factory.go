package factoryperson

import (
	"sync"

	dbcommon "github.com/darllantissei/genealogy-tree/adapters/db/sqlite/common"
	dbconnectionrelationship "github.com/darllantissei/genealogy-tree/adapters/db/sqlite/relationship"
	servicecommon "github.com/darllantissei/genealogy-tree/application/common"
	personservice "github.com/darllantissei/genealogy-tree/application/person"
	servicerelationship "github.com/darllantissei/genealogy-tree/application/relationship"
)

type FactoryRelationship struct {
	CommonDB            dbcommon.CommonDB
	CommonService       servicecommon.ICommonService
	PersonService       personservice.IPersonService
	serviceRelationship *servicerelationship.RelationshipService
	dbconnection        *dbconnectionrelationship.RelationshipDB
	lockFlow            *sync.Mutex
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
		fr.dbconnection = dbconnectionrelationship.NewRelationshipDB(fr.CommonDB)
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
			CommonService: fr.CommonService,
			PersonService: fr.PersonService,
		}

	}

	return fr.serviceRelationship
}
