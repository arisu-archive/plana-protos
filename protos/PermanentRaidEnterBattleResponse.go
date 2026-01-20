package protos

type PermanentRaidEnterBattleResponse struct {
	ResponsePacket
	BattleHistoryDB PermanentRaidBattleHistoryDB `json:",omitempty,omitzero"`
}
