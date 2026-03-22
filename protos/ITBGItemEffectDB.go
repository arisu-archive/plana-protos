package protos

type ITBGItemEffectDB struct {
	ItemInfo               ITBGItemInfo
	Stack                  int32 `json:",omitempty,omitzero"`
	RemainEncounterCounter int32 `json:",omitempty,omitzero"`
}
