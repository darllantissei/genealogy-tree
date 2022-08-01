package model

import enumgender "github.com/darllantissei/genealogy-tree/application/enum/gender"

type Person struct {
	ID        string            `valid:"-"`
	FirstName string            `valid:"required~The First name is required"`
	LastName  string            `valid:"optional"`
	Gender    enumgender.Gender `valid:"optional"`
}
