package protos

type EventRewardIncreaseResponse struct {
	ResponsePacket
	EventRewardIncreaseDBs []*EventRewardIncreaseDB
}
