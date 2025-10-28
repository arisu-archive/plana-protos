package protos

type ClanAllAssistListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AssistCharacterDBs []AssistCharacterDB `json:",omitempty,omitzero"`
	AssistCharacterRentHistoryDBs []ClanAssistRentHistoryDB `json:",omitempty,omitzero"`
	ClanDBId int64 `json:",omitempty,omitzero"`
}
