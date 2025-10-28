package protos

type RaidDetailResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidDetailDB RaidDetailDB `json:",omitempty,omitzero"`
	ParticipateCharacterServerIds []int64 `json:",omitempty,omitzero"`
}
