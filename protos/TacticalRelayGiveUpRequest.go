package protos

type TacticalRelayGiveUpRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
	StageId  int64 `json:",omitempty,omitzero"`
}
