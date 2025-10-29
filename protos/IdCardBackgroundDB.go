package protos

type IdCardBackgroundDB struct {
	ParcelBase
	ServerId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
}
