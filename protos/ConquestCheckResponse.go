package protos

type ConquestCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CanReceiveCalculateReward bool `json:",omitempty,omitzero"`
	AlarmPhaseToShow *int32 `json:",omitempty,omitzero"`
	ParcelConsumeCumulatedAmount int64 `json:",omitempty,omitzero"`
	ConquestSummary ConquestSummary `json:",omitempty,omitzero"`
}
