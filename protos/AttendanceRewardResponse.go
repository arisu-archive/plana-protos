package protos

type AttendanceRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AttendanceBookRewards []AttendanceBookReward `json:",omitempty,omitzero"`
	AttendanceHistoryDBs []AttendanceHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
