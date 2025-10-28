package protos

type EventContentLocationAttendScheduleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentLocationDB EventContentLocationDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ExtraRewards []ParcelInfo `json:",omitempty,omitzero"`
}
