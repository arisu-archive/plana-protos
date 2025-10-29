package protos

type EventContentLocationAttendScheduleResponse struct {
	ResponsePacket
	EventContentLocationDB EventContentLocationDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ExtraRewards []ParcelInfo `json:",omitempty,omitzero"`
}
