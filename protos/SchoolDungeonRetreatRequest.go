package protos

type SchoolDungeonRetreatRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
