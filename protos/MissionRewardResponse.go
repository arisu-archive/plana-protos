package protos

type MissionRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AddedHistoryDB MissionHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
