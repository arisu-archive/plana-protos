package protos

type TBGEncounterDB struct {
	EncounterId                     int64           `json:",omitempty,omitzero"`
	InvokerServerId                 int64           `json:",omitempty,omitzero"`
	ObjectType                      int32           `json:",omitempty,omitzero"`
	ShouldDecreaseItemEffectCounter bool            `json:",omitempty,omitzero"`
	RewardUniqueIdByIndex           map[int32]int64 `json:",omitempty,omitzero"`
}
