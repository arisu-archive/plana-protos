package protos

type RelaySquadRecord struct {
	SquadIndex            int32 `json:",omitempty,omitzero"`
	StrikerServerIds      []int64
	SupporterServerIds    []int64
	ApcStrikerServerIds   []int64
	ApcSupporterServerIds []int64
}
