package protos

import (
	"time"
)

type IssueAlertInfoDB struct {
	IssueAlertId int32 `json:",omitempty,omitzero"`
	IssueAlertType IssueAlertTypeCode `json:",omitempty,omitzero"`
	StartDate time.Time `json:",omitempty,omitzero"`
	EndDate time.Time `json:",omitempty,omitzero"`
	DisplayOrder byte `json:",omitempty,omitzero"`
	PublishId int32 `json:",omitempty,omitzero"`
	Url string `json:",omitempty,omitzero"`
	Subject string `json:",omitempty,omitzero"`
}
