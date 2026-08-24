package protos

type AccountGachaWishDB struct {
	SelectAbleGachaGroupId int64 `json:",omitempty,omitzero"`
	SelectedCharacterIds   []int64
}
