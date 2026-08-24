package protos

type StudentFrameDB struct {
	ParcelBase
	UniqueId    int64  `json:",omitempty,omitzero"`
	ReceiveDate MxTime `json:",omitempty,omitzero"`
}
