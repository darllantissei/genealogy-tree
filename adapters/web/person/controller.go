package personweb

import (
	"github.com/darllantissei/genealogy-tree/application/model"
)

func (pw *PersonWeb) parseDataToApp(personDTO PersonDTO) model.Person {

	return model.Person{
		ID:        personDTO.ID,
		FirstName: personDTO.FirstName,
		LastName:  personDTO.LastName,
		Gender:    personDTO.Gender,
	}
}

func (pw *PersonWeb) parseData(personApp model.Person) PersonDTO {

	return PersonDTO{
		ID:        personApp.ID,
		FirstName: personApp.FirstName,
		LastName:  personApp.LastName,
		Gender:    personApp.Gender,
	}
}
