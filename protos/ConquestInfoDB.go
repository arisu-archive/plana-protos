package protos

type ConquestInfoDB struct {
	EventContentId                         int64  `json:",omitempty,omitzero"`
	EventGauge                             int32  `json:",omitempty,omitzero"`
	EventSpawnCount                        int32  `json:",omitempty,omitzero"`
	EchelonChangeCount                     int32  `json:",omitempty,omitzero"`
	TodayConquestRentCount                 int32  `json:",omitempty,omitzero"`
	TodayOperationRentCount                int32  `json:",omitempty,omitzero"`
	CumulatedConditionValue                int64  `json:",omitempty,omitzero"`
	ReceivedCalculateRewardConditionAmount int64  `json:",omitempty,omitzero"`
	AlertMassErosionId                     *int64 `json:",omitempty,omitzero"`
}
