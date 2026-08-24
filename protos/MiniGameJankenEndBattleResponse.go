package protos

type MiniGameJankenEndBattleResponse struct {
	ResponsePacket
	SaveDB                     *MiniGameJankenSaveDB         `json:",omitempty,omitzero"`
	StageHistoryDB             *MiniGameJankenStageHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB             *ParcelResultDB               `json:",omitempty,omitzero"`
	FirstClearRewards          []*ParcelInfo
	FirstThreeStarClearRewards []*ParcelInfo
}
