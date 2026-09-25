package protos

type CharacterAdaptationGrowthStepResponse struct {
	ResponsePacket
	CharacterAdaptationDB *CharacterAdaptationDB `json:",omitempty,omitzero"`
}
