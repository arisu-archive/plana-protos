package protos

type NotificationLobbyCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UnreadMailCount int64 `json:",omitempty,omitzero"`
	EventRewardIncreaseDBs []EventRewardIncreaseDB `json:",omitempty,omitzero"`
}
