package protos

type RaidGiveUpResponse struct {
	ResponsePacket
	Tier           int32 `json:",omitempty,omitzero"`
	RaidGiveUpDB   RaidGiveUpDB
	ParcelResultDB ParcelResultDB
}
