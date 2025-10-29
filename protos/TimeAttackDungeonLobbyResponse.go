package protos

type TimeAttackDungeonLobbyResponse struct {
	ResponsePacket
	RoomDBs map[int64]TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
	PreviousRoomDB TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	AchieveSeasonBestRecord bool `json:",omitempty,omitzero"`
	SeasonBestRecord int64 `json:",omitempty,omitzero"`
}
