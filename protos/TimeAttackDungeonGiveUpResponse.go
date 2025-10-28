package protos

type TimeAttackDungeonGiveUpResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RoomDB TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	AchieveSeasonBestRecord bool `json:",omitempty,omitzero"`
	SeasonBestRecord int64 `json:",omitempty,omitzero"`
}
