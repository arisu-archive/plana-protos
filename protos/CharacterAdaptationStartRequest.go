package protos

type CharacterAdaptationStartRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
