package protos

type CharacterAdaptationStartResponse struct {
	ResponsePacket
	CharacterAdaptationDB *CharacterAdaptationDB `json:",omitempty,omitzero"`
}
