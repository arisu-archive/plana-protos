package protos

type AcademyAttendScheduleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	AcademyDB AcademyDB `json:",omitempty,omitzero"`
	ExtraRewards []ParcelInfo `json:",omitempty,omitzero"`
}
