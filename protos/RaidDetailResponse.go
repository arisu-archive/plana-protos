package protos

type RaidDetailResponse struct {
	ResponsePacket
	RaidDetailDB                  RaidDetailDB
	ParticipateCharacterServerIds []int64
}
