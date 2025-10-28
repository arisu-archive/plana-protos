package protos

type WorldRaidEnterBattleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidBattleDB RaidBattleDB `json:",omitempty,omitzero"`
	AssistCharacterDBs []AssistCharacterDB `json:",omitempty,omitzero"`
}
