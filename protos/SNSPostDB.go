package protos

type SNSPostDB struct {
	ParcelBase
	SNSId       int64   `json:",omitempty,omitzero"`
	PostId      int64   `json:",omitempty,omitzero"`
	ReceiveDate MxTime  `json:",omitempty,omitzero"`
	ReadDate    *MxTime `json:",omitempty,omitzero"`
}
