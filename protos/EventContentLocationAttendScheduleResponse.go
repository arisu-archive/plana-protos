package protos

type EventContentLocationAttendScheduleResponse struct {
	ResponsePacket
	EventContentLocationDB    *EventContentLocationDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []*EventContentCollectionDB
	ParcelResultDB            *ParcelResultDB `json:",omitempty,omitzero"`
	ExtraRewards              []*ParcelInfo
}
