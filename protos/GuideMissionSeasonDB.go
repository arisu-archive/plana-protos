package protos

type GuideMissionSeasonDB struct {
	SeasonId                  int64   `json:",omitempty,omitzero"`
	LoginCount                int64   `json:",omitempty,omitzero"`
	StartDate                 MxTime  `json:",omitempty,omitzero"`
	LoginDate                 MxTime  `json:",omitempty,omitzero"`
	IsComplete                bool    `json:",omitempty,omitzero"`
	IsFinalMissionComplete    bool    `json:",omitempty,omitzero"`
	CollectionItemReceiveDate *MxTime `json:",omitempty,omitzero"`
}
