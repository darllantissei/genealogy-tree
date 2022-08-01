package model

import (
	"strings"

	enumstatusapp "github.com/darllantissei/genealogy-tree/application/enum/status_application"
)

type Returns struct {
	Return `json:"return" valid:"-"`
}

type Return struct {
	Status  enumstatusapp.StatusApp `json:"status" valid:"-"`
	Message []string `json:"message,omitempty" valid:"-"`
}

func (r Returns) Error() string {
	return strings.Join(r.Return.Message, ";")
}
