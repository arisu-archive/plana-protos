package protos

type TimeAttackDungeonGiveUpResponse struct {
	ResponsePacket
	RoomDB                  TimeAttackDungeonRoomDB
	ParcelResultDB          ParcelResultDB
	AchieveSeasonBestRecord bool  `json:",omitempty,omitzero"`
	SeasonBestRecord        int64 `json:",omitempty,omitzero"`
}
