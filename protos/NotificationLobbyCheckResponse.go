package protos

type NotificationLobbyCheckResponse struct {
	ResponsePacket
	UnreadMailCount              int64 `json:",omitempty,omitzero"`
	UnreadCommonMailCount        int64 `json:",omitempty,omitzero"`
	UnreadSemiPermanentMailCount int64 `json:",omitempty,omitzero"`
	EventRewardIncreaseDBs       []EventRewardIncreaseDB
}
