package protos

type AcademyGetInfoResponse struct {
	ResponsePacket
	AcademyDB          AcademyDB           `json:",omitempty,omitzero"`
	AcademyLocationDBs []AcademyLocationDB `json:",omitempty,omitzero"`
}
