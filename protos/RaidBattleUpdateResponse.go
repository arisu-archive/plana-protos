package protos

type RaidBattleUpdateResponse struct {
	ResponsePacket
	RaidBattleDB RaidBattleDB `json:",omitempty,omitzero"`
}
