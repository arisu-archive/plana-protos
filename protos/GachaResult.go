package protos

type GachaResult struct {
	CharacterId int64 `json:",omitempty,omitzero"`
	Character   CharacterDB
	Stone       ItemDB
}
