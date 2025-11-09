package protos

type CampaignSubStageResultRequest struct {
	RequestPacket
	PassCheckCharacter bool          `json:",omitempty,omitzero"`
	Summary            BattleSummary `json:",omitempty,omitzero"`
}
