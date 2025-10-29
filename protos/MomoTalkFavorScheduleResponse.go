package protos

type MomoTalkFavorScheduleResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	FavorScheduleRecords map[int64][]int64 `json:",omitempty,omitzero"`
}
