package protos

type CharacterAdaptationDB struct {
	SeasonId    int64 `json:",omitempty,omitzero"`
	CharacterId int64 `json:",omitempty,omitzero"`
	MissionStep int32 `json:",omitempty,omitzero"`
	IsComplete  bool  `json:",omitempty,omitzero"`
}
