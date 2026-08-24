package protos

type MiniGameJankenScoreRewardRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
