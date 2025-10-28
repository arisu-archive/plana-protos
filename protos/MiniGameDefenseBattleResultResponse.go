package protos

type MiniGameDefenseBattleResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	StageHistoryDB MiniGameDefenseStageHistoryDB `json:",omitempty,omitzero"`
}
