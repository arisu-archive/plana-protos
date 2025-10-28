package protos

type AccountRequestBirthdayMailResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
