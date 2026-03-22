package protos

type EventContentLocationAttendScheduleResponse struct {
	ResponsePacket
	EventContentLocationDB    EventContentLocationDB
	EventContentCollectionDBs []EventContentCollectionDB
	ParcelResultDB            ParcelResultDB
	ExtraRewards              []ParcelInfo
}
