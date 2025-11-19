package protos

type AcademyAttendScheduleResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	AcademyDB      AcademyDB      `json:",omitempty,omitzero"`
	ExtraRewards   []ParcelInfo   `json:",omitempty,omitzero"`
}
