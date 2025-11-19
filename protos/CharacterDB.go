package protos

type CharacterDB struct {
	ParcelBase
	ServerId               int64           `json:",omitempty,omitzero"`
	UniqueId               int64           `json:",omitempty,omitzero"`
	StarGrade              int32           `json:",omitempty,omitzero"`
	Level                  int32           `json:",omitempty,omitzero"`
	Exp                    int64           `json:",omitempty,omitzero"`
	FavorRank              int32           `json:",omitempty,omitzero"`
	FavorExp               int64           `json:",omitempty,omitzero"`
	PublicSkillLevel       int32           `json:",omitempty,omitzero"`
	ExSkillLevel           int32           `json:",omitempty,omitzero"`
	PassiveSkillLevel      int32           `json:",omitempty,omitzero"`
	ExtraPassiveSkillLevel int32           `json:",omitempty,omitzero"`
	LeaderSkillLevel       int32           `json:",omitempty,omitzero"`
	IsFavorite             bool            `json:",omitempty,omitzero"`
	EquipmentServerIds     []int64         `json:",omitempty,omitzero"`
	PotentialStats         map[int32]int32 `json:",omitempty,omitzero"`
}
