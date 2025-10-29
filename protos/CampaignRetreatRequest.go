package protos

type CampaignRetreatRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
