package protos

type MiniGameCCGSelectRewardCardRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	SelectedIndex  int32 `json:",omitempty,omitzero"`
	RewardIndex    int32 `json:",omitempty,omitzero"`
}
