package protos

type ArenaDailyRewardResponse struct {
	ResponsePacket
	ParcelResult          ParcelResultDB `json:",omitempty,omitzero"`
	DailyRewardActiveTime MxTime         `json:",omitempty,omitzero"`
}
