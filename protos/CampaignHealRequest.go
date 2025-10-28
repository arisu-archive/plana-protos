package protos

type CampaignHealRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CampaignStageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonIndex int64 `json:",omitempty,omitzero"`
	CharacterServerId int64 `json:",omitempty,omitzero"`
}
