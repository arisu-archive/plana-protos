package protos

type MiniGameTankRaceLevelUpRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	TankGroupId    int64 `json:",omitempty,omitzero"`
}
