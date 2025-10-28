package protos

type UseCouponResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CouponCompleteRewardReceived bool `json:",omitempty,omitzero"`
}
