package protos

type IssueAlertInfoDB struct {
	IssueAlertId   int32              `json:",omitempty,omitzero"`
	IssueAlertType IssueAlertTypeCode `json:",omitempty,omitzero"`
	StartDate      MxTime             `json:",omitempty,omitzero"`
	EndDate        MxTime             `json:",omitempty,omitzero"`
	DisplayOrder   byte               `json:",omitempty,omitzero"`
	PublishId      int32              `json:",omitempty,omitzero"`
	Url            string             `json:",omitempty,omitzero"`
	Subject        string             `json:",omitempty,omitzero"`
}
