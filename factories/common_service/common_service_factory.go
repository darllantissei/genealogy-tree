package factoryperson

import (
	"sync"

	servicecommon "github.com/darllantissei/genealogy-tree/application/common"
)

type FactoryCommonService struct {
	serviceCommon *servicecommon.CommonService
	lockFlow      *sync.Mutex
}

func (fp *FactoryCommonService) NewService() *servicecommon.CommonService {

	if fp.lockFlow == nil {

		fp.lockFlow = &sync.Mutex{}

		fp.lockFlow.Lock()

		defer func() {

			fp.lockFlow.Unlock()

			fp.lockFlow = nil
		}()

	}

	if fp.serviceCommon == nil {

		fp.serviceCommon = &servicecommon.CommonService{}

	}

	return fp.serviceCommon
}
