package protos

type UseCouponRequest struct {
	RequestPacket
	CouponSerial string `json:",omitempty,omitzero"`
}
