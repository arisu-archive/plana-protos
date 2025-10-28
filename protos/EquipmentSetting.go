package protos

type EquipmentSetting struct {
	ServerId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
	Tier int32 `json:",omitempty,omitzero"`
}
