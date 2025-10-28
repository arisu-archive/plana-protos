package protos

type MomoTalkFavorScheduleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	FavorScheduleRecords map[int64][]int64 `json:",omitempty,omitzero"`
}
