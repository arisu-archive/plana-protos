package protos

func (r *ResponsePacket) GetResponsePacket() *ResponsePacket {
	return r
}

func (r *ResponsePacket) SetResponsePacket(nr ResponsePacket) {
	*r = nr
}
