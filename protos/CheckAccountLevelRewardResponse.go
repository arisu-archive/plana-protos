package protos

type CheckAccountLevelRewardResponse struct {
	ResponsePacket
	AccountLevelRewardIds []int64 `json:",omitempty,omitzero"`
}
