package model

import (
	"strings"

	enumstatusapp "github.com/darllantissei/genealogy-tree/application/enum/status_application"
)

type Returns struct {
	XMLName struct{} `json:"-" xml:"return"`
	Return  `json:"return" xml:"return" valid:"-"`
}

type Return struct {
	Status  enumstatusapp.StatusApp `json:"status" xml:"status" valid:"-"`
	Message []string                `json:"message,omitempty" xml:"message,omitempty" valid:"-"`
}

func (r Returns) Error() string {
	return strings.Join(r.Return.Message, ";")
}
