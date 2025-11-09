package protos

type AttendanceRewardResponse struct {
	ResponsePacket
	AttendanceBookRewards []AttendanceBookReward `json:",omitempty,omitzero"`
	AttendanceHistoryDBs  []AttendanceHistoryDB  `json:",omitempty,omitzero"`
	ParcelResultDB        ParcelResultDB         `json:",omitempty,omitzero"`
}
