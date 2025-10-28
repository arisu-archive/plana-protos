package protos

type EventRewardIncreaseResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventRewardIncreaseDBs []EventRewardIncreaseDB `json:",omitempty,omitzero"`
}
