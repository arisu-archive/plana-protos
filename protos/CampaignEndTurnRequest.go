package protos

type CampaignEndTurnRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
