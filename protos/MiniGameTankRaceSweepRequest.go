package protos

type MiniGameTankRaceSweepRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageId        int64 `json:",omitempty,omitzero"`
	SweepCount     int64 `json:",omitempty,omitzero"`
}
