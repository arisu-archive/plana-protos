package protos

type CharacterExpGrowthResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CharacterDB CharacterDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
