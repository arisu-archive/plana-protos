package protos

type UseCouponRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CouponSerial string `json:",omitempty,omitzero"`
}
