package protos

type ArenaCumulativeTimeRewardResponse struct {
	ResponsePacket
	TimeRewardAmount         int64          `json:",omitempty,omitzero"`
	TimeRewardLastUpdateTime MxTime         `json:",omitempty,omitzero"`
	ParcelResult             ParcelResultDB `json:",omitempty,omitzero"`
}
