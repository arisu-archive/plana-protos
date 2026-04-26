package protos

type SNSPostReadRequest struct {
	RequestPacket
	SNSId int64 `json:",omitempty,omitzero"`
}
