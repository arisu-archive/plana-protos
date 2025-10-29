package protos

type WeaponDB struct {
	ParcelBase
	UniqueId int64 `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
	Exp int64 `json:",omitempty,omitzero"`
	StarGrade int32 `json:",omitempty,omitzero"`
	BoundCharacterServerId int64 `json:",omitempty,omitzero"`
}
