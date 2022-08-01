package person

import (
	"context"

	"github.com/darllantissei/genealogy-tree/application/common"
	"github.com/darllantissei/genealogy-tree/application/model"
)

func (p *PersonService) checkDataPerson(ctx context.Context, prsn model.Person) (msgErr []string) {

	_, err := common.Validation(prsn)

	if err != nil {
		msgErr = append(msgErr, err.Error())
	}

	err = prsn.Gender.Scan(prsn.Gender)

	if err != nil {
		msgErr = append(msgErr, err.Error())
	}

	return msgErr
}
