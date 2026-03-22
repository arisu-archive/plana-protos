package protos

type EliminateRaidGiveUpResponse struct {
	ResponsePacket
	Tier           int32 `json:",omitempty,omitzero"`
	RaidGiveUpDB   RaidGiveUpDB
	ParcelResultDB ParcelResultDB
}
