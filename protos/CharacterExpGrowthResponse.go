package protos

type CharacterExpGrowthResponse struct {
	ResponsePacket
	CharacterDB       CharacterDB       `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ConsumeResultDB   ConsumeResultDB   `json:",omitempty,omitzero"`
}
