package protos

type MissionRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MissionUniqueId int64 `json:",omitempty,omitzero"`
	ProgressServerId int64 `json:",omitempty,omitzero"`
	EventContentId *int64 `json:",omitempty,omitzero"`
}
