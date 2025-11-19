package protos

type EventContentLocationDB struct {
	LocationId            int64                           `json:",omitempty,omitzero"`
	Rank                  int64                           `json:",omitempty,omitzero"`
	Exp                   int64                           `json:",omitempty,omitzero"`
	ScheduleCount         int64                           `json:",omitempty,omitzero"`
	ZoneVisitCharacterDBs map[int64][]VisitingCharacterDB `json:",omitempty,omitzero"`
}
