package protos

type DailyRecordRewardResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	DailyRecordDB  DailyRecordDB
}
