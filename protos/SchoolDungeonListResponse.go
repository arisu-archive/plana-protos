package protos

type SchoolDungeonListResponse struct {
	ResponsePacket
	SchoolDungeonStageHistoryDBList []SchoolDungeonStageHistoryDB `json:",omitempty,omitzero"`
}
