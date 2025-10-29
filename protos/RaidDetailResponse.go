package protos

type RaidDetailResponse struct {
	ResponsePacket
	RaidDetailDB RaidDetailDB `json:",omitempty,omitzero"`
	ParticipateCharacterServerIds []int64 `json:",omitempty,omitzero"`
}
