package protos

type PermanentRaidBestScoreHistoryDB struct {
	StageId      int64         `json:",omitempty,omitzero"`
	Score        int64         `json:",omitempty,omitzero"`
	ClearDeckDBs []ClearDeckDB `json:",omitempty,omitzero"`
}
