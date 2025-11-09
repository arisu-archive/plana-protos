package protos

type WorldRaidEnterBattleResponse struct {
	ResponsePacket
	RaidBattleDB       RaidBattleDB        `json:",omitempty,omitzero"`
	AssistCharacterDBs []AssistCharacterDB `json:",omitempty,omitzero"`
}
