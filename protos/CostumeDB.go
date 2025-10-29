package protos

type CostumeDB struct {
	ParcelBase
	UniqueId int64 `json:",omitempty,omitzero"`
	BoundCharacterServerId int64 `json:",omitempty,omitzero"`
}
