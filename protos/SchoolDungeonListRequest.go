package protos

type SchoolDungeonListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
