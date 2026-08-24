package protos

type DailyRecordClaimSeasonRewardResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB           `json:",omitempty,omitzero"`
	SeasonRecordDB *AttendanceSeasonRecordDB `json:",omitempty,omitzero"`
}
