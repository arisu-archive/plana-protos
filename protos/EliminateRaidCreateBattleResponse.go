package protos

type EliminateRaidCreateBattleResponse struct {
	ResponsePacket
	RaidDB            RaidDB            `json:",omitempty,omitzero"`
	RaidBattleDB      RaidBattleDB      `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	AssistCharacterDB AssistCharacterDB `json:",omitempty,omitzero"`
}
