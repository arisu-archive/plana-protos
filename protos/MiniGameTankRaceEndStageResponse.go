package protos

type MiniGameTankRaceEndStageResponse struct {
	ResponsePacket
	StageHistoryDB    *MiniGameTankRaceHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB    *ParcelResultDB            `json:",omitempty,omitzero"`
	FirstClearRewards []*ParcelInfo
	ThreeStarRewards  []*ParcelInfo
}
