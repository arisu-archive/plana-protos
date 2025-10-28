package protos

type RaidBattleUpdateResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidBattleDB RaidBattleDB `json:",omitempty,omitzero"`
}
