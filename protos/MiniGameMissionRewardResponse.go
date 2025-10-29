package protos

type MiniGameMissionRewardResponse struct {
	ResponsePacket
	AddedHistoryDB MissionHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
