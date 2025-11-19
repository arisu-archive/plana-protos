package protos

type CampaignPortalRequest struct {
	RequestPacket
	StageUniqueId   int64 `json:",omitempty,omitzero"`
	EchelonEntityId int64 `json:",omitempty,omitzero"`
}
