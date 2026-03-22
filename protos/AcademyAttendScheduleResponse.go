package protos

type AcademyAttendScheduleResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	AcademyDB      AcademyDB
	ExtraRewards   []ParcelInfo
}
