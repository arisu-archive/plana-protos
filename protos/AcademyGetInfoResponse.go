package protos

type AcademyGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AcademyDB AcademyDB `json:",omitempty,omitzero"`
	AcademyLocationDBs []AcademyLocationDB `json:",omitempty,omitzero"`
}
