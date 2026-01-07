package protos

type ITBGItemEffectDB struct {
	ItemInfo               ITBGItemInfo `json:",omitempty,omitzero"`
	Stack                  int32        `json:",omitempty,omitzero"`
	RemainEncounterCounter int32        `json:",omitempty,omitzero"`
}
