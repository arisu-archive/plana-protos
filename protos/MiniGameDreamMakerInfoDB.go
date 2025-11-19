package protos

type MiniGameDreamMakerInfoDB struct {
	EventContentId       int64 `json:",omitempty,omitzero"`
	Round                int64 `json:",omitempty,omitzero"`
	Multiplier           int64 `json:",omitempty,omitzero"`
	DayOfNumber          int64 `json:",omitempty,omitzero"`
	ActionCount          int64 `json:",omitempty,omitzero"`
	CurrentRoundEndingId int64 `json:",omitempty,omitzero"`
	EndingRewardReceived bool  `json:",omitempty,omitzero"`
	CanStartNewGame      bool  `json:",omitempty,omitzero"`
}
