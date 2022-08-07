package model

import (
	"fmt"
	"reflect"

	enumgender "github.com/darllantissei/genealogy-tree/application/enum/gender"
)

type Person struct {
	ID        string            `valid:"-"`
	FirstName string            `valid:"required~The First name is required"`
	LastName  string            `valid:"optional"`
	Gender    enumgender.Gender `valid:"optional"`
}

func (prsn *Person) IsEmpty() bool {

	if prsn == nil {
		return true
	}

	if reflect.DeepEqual(prsn, &Person{}) {
		return true
	}

	return false
}

func (prsn Person) FullName() string {
	return fmt.Sprintf("%s %s", prsn.FirstName, prsn.LastName)
}
