package protos

type GearSetting struct {
	UniqueId int64 `json:",omitempty,omitzero"`
	Tier int32 `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
}
