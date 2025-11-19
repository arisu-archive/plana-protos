package protos

type ConsumableItemBaseDB struct {
	ParcelBase
	ServerId   int64 `json:",omitempty,omitzero"`
	UniqueId   int64 `json:",omitempty,omitzero"`
	StackCount int64 `json:",omitempty,omitzero"`
}
