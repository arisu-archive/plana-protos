package protos

type WeaponSetting struct {
	UniqueId  int64 `json:",omitempty,omitzero"`
	StarGrade int32 `json:",omitempty,omitzero"`
	Level     int32 `json:",omitempty,omitzero"`
}
