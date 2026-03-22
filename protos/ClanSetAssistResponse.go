package protos

type ClanSetAssistResponse struct {
	ResponsePacket
	ClanAssistSlotDB ClanAssistSlotDB
	ParcelResultDB   ParcelResultDB
	RewardInfo       ClanAssistRewardInfo
}
