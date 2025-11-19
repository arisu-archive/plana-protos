package protos

type CheatCharacterCustomPreset struct {
	UniqueId            int64                        `json:",omitempty,omitzero"`
	StarGrade           int32                        `json:",omitempty,omitzero"`
	Level               int32                        `json:",omitempty,omitzero"`
	ExSkillLevel        int32                        `json:",omitempty,omitzero"`
	PublicSkillLevel    int32                        `json:",omitempty,omitzero"`
	PassiveSkillLevel   int32                        `json:",omitempty,omitzero"`
	ExPassiveSkillLevel int32                        `json:",omitempty,omitzero"`
	Equipments          []CheatEquipmentCustomPreset `json:",omitempty,omitzero"`
	Weapon              CheatWeaponCustomPreset      `json:",omitempty,omitzero"`
}
