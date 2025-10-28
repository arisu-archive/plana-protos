package protos

type CampaignWithdrawEchelonRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	WithdrawEchelonEntityId []int64 `json:",omitempty,omitzero"`
}
