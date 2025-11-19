package protos

type NotificationLobbyCheckResponse struct {
	ResponsePacket
	UnreadMailCount        int64                   `json:",omitempty,omitzero"`
	EventRewardIncreaseDBs []EventRewardIncreaseDB `json:",omitempty,omitzero"`
}
