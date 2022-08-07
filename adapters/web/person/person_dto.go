package personweb

import (
	"reflect"

	enumgender "github.com/darllantissei/genealogy-tree/application/enum/gender"
)

type PersonDTO struct {
	XMLName   struct{}          `json:"-" xml:"person"`
	ID        string            `json:"id" xml:"id"`
	FirstName string            `json:"first_name" xml:"first_name"`
	LastName  string            `json:"last_name" xml:"last_name"`
	Gender    enumgender.Gender `json:"gender" xml:"gender"`
}

func (pdto *PersonDTO) IsEmpty() bool {

	if pdto == nil {
		return true
	}

	if reflect.DeepEqual(pdto, &PersonDTO{}) {
		return true
	}

	return false
}
