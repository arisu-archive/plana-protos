package protos

type UseCouponResponse struct {
	ResponsePacket
	CouponCompleteRewardReceived bool `json:",omitempty,omitzero"`
}
