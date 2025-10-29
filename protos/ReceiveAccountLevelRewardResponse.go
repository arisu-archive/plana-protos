package protos

type ReceiveAccountLevelRewardResponse struct {
	ResponsePacket
	ReceivedAccountLevelRewardIds []int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
