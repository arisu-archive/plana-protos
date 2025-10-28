package protos

type RaidGiveUpResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Tier int32 `json:",omitempty,omitzero"`
	RaidGiveUpDB RaidGiveUpDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
