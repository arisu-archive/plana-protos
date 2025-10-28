package protos

type SchoolDungeonListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SchoolDungeonStageHistoryDBList []SchoolDungeonStageHistoryDB `json:",omitempty,omitzero"`
}
