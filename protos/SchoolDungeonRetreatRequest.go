package protos

type SchoolDungeonRetreatRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
