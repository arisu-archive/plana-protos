package protos

type AcademyGetInfoRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
