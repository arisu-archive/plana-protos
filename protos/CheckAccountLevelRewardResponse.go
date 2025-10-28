package protos

type CheckAccountLevelRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountLevelRewardIds []int64 `json:",omitempty,omitzero"`
}
