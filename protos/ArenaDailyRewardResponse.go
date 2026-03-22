package protos

type ArenaDailyRewardResponse struct {
	ResponsePacket
	ParcelResult          ParcelResultDB
	DailyRewardActiveTime MxTime `json:",omitempty,omitzero"`
}
