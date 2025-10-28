package protos

type EliminateRaidEnterBattleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidDB RaidDB `json:",omitempty,omitzero"`
	RaidBattleDB RaidBattleDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	AssistCharacterDB AssistCharacterDB `json:",omitempty,omitzero"`
}
