package protos

type MiniGameTankRaceEndFreeStageRequest struct {
	RequestPacket
	EventContentId int64                    `json:",omitempty,omitzero"`
	StageId        int64                    `json:",omitempty,omitzero"`
	Summary        *MiniGameTankRaceSummary `json:",omitempty,omitzero"`
}
