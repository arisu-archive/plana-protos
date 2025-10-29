package protos

type CampaignHealRequest struct {
	RequestPacket
	CampaignStageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonIndex int64 `json:",omitempty,omitzero"`
	CharacterServerId int64 `json:",omitempty,omitzero"`
}
