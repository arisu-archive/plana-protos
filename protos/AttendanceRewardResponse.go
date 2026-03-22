package protos

type AttendanceRewardResponse struct {
	ResponsePacket
	AttendanceBookRewards []AttendanceBookReward
	AttendanceHistoryDBs  []AttendanceHistoryDB
	ParcelResultDB        ParcelResultDB
}
