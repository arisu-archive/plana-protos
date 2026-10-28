package protos

type MiniGameTankRaceEndFreeStageResponse struct {
	ResponsePacket
	StageHistoryDB *MiniGameTankRaceHistoryDB `json:",omitempty,omitzero"`
}
