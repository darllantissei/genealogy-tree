package factoryperson

import (
	"sync"

	dbcommon "github.com/darllantissei/genealogy-tree/adapters/db/sqlite/common"
	dbconnectionperson "github.com/darllantissei/genealogy-tree/adapters/db/sqlite/person"
	servicecommon "github.com/darllantissei/genealogy-tree/application/common"
	serviceperson "github.com/darllantissei/genealogy-tree/application/person"
)

type FactoryPerson struct {
	CommonDB      dbcommon.CommonDB
	CommonService servicecommon.ICommonService
	servicePerson *serviceperson.PersonService
	dbconnection  *dbconnectionperson.PersonDB
	lockFlow      *sync.Mutex
}

func (fp *FactoryPerson) NewDataBaseConnection() *dbconnectionperson.PersonDB {

	if fp.lockFlow == nil {

		fp.lockFlow = &sync.Mutex{}

		fp.lockFlow.Lock()

		defer func() {

			fp.lockFlow.Unlock()

			fp.lockFlow = nil
		}()

	}

	if fp.dbconnection == nil {
		fp.dbconnection = dbconnectionperson.NewPersonDB(fp.CommonDB)
	}

	return fp.dbconnection
}

func (fp *FactoryPerson) NewService() *serviceperson.PersonService {

	if fp.lockFlow == nil {

		fp.lockFlow = &sync.Mutex{}

		fp.lockFlow.Lock()

		defer func() {

			fp.lockFlow.Unlock()

			fp.lockFlow = nil
		}()

	}

	if fp.servicePerson == nil {

		fp.servicePerson = &serviceperson.PersonService{
			PersistenceDB: fp.NewDataBaseConnection(),
			CommonService: fp.CommonService,
		}

	}

	return fp.servicePerson
}
