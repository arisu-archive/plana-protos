package protos

type AccountRequestBirthdayMailRequest struct {
	RequestPacket
	Birthday MxTime `json:",omitempty,omitzero"`
}
