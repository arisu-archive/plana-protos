package protos

type MiniGameTankRaceHistoryDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	StageId        int64 `json:",omitempty,omitzero"`
	Star1Flag      bool  `json:",omitempty,omitzero"`
	Star2Flag      bool  `json:",omitempty,omitzero"`
	Star3Flag      bool  `json:",omitempty,omitzero"`
	BestRecord     int64 `json:",omitempty,omitzero"`
	StarCount      int64 `json:",omitempty,omitzero"`
}
