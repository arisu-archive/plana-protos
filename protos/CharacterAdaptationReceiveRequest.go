package protos

type CharacterAdaptationReceiveRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
