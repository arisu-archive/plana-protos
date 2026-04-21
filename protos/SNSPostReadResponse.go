package protos

type SNSPostReadResponse struct {
	ResponsePacket
	SNSPostDBs []*SNSPostDB
}
