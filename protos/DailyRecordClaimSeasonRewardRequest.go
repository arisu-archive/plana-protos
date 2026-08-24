package protos

type DailyRecordClaimSeasonRewardRequest struct {
	RequestPacket
	SeasonRecordId int64 `json:",omitempty,omitzero"`
}
