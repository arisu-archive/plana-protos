package protos

type EventContentConcentrationGetInfoRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
