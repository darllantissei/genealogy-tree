package common

import (
	"fmt"
	"reflect"

	gouuid "github.com/satori/go.uuid"
)

func (c CommonService) SliceExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic(fmt.Sprintf("SliceExists(). The type %v is not a type slice valid", s.Kind()))
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func (c CommonService) GetUUID() string {

	uuid := gouuid.NewV4()

	return uuid.String()
}
