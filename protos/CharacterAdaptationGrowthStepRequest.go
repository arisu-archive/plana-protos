package protos

type CharacterAdaptationGrowthStepRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
