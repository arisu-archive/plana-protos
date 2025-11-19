package protos

type GachaResult struct {
	CharacterId int64       `json:",omitempty,omitzero"`
	Character   CharacterDB `json:",omitempty,omitzero"`
	Stone       ItemDB      `json:",omitempty,omitzero"`
}
