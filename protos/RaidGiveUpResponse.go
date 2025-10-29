package protos

type RaidGiveUpResponse struct {
	ResponsePacket
	Tier int32 `json:",omitempty,omitzero"`
	RaidGiveUpDB RaidGiveUpDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
