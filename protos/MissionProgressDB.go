package protos

type MissionProgressDB struct {
	MissionUniqueId    int64           `json:",omitempty,omitzero"`
	Complete           bool            `json:",omitempty,omitzero"`
	StartTime          MxTime          `json:",omitempty,omitzero"`
	ProgressParameters map[int64]int64 `json:",omitempty,omitzero"`
}
