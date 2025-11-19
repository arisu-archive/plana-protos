package protos

type MiniGameDefenseBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB                `json:",omitempty,omitzero"`
	StageHistoryDB MiniGameDefenseStageHistoryDB `json:",omitempty,omitzero"`
}
