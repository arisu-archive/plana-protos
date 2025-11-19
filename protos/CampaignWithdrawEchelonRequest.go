package protos

type CampaignWithdrawEchelonRequest struct {
	RequestPacket
	StageUniqueId           int64   `json:",omitempty,omitzero"`
	WithdrawEchelonEntityId []int64 `json:",omitempty,omitzero"`
}
