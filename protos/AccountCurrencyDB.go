package protos

type AccountCurrencyDB struct {
	AccountLevel           int64             `json:",omitempty,omitzero"`
	AcademyLocationRankSum int64             `json:",omitempty,omitzero"`
	CurrencyDict           map[string]int64  `json:",omitempty,omitzero"`
	UpdateTimeDict         map[string]MxTime `json:",omitempty,omitzero"`
}
