package protos

type CampaignPortalRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
}
