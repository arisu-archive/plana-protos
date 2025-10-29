package protos

type EventContentDiceRaceLapRewardRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
