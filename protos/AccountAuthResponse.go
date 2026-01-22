package protos

type AccountAuthResponse struct {
	ResponsePacket
	CurrentVersion                      int64                  `json:",omitempty,omitzero"`
	MinimumVersion                      int64                  `json:",omitempty,omitzero"`
	IsDevelopment                       bool                   `json:",omitempty,omitzero"`
	BattleValidation                    bool                   `json:",omitempty,omitzero"`
	UpdateRequired                      bool                   `json:",omitempty,omitzero"`
	TTSCdnUri                           string                 `json:",omitempty,omitzero"`
	AccountDB                           AccountDB              `json:",omitempty,omitzero"`
	AttendanceBookRewards               []AttendanceBookReward `json:",omitempty,omitzero"`
	AttendanceHistoryDBs                []AttendanceHistoryDB  `json:",omitempty,omitzero"`
	RepurchasableMonthlyProductCountDBs []PurchaseCountDB      `json:",omitempty,omitzero"`
	MonthlyProductParcel                []ParcelInfo           `json:",omitempty,omitzero"`
	MonthlyProductMail                  []ParcelInfo           `json:",omitempty,omitzero"`
	BiweeklyProductParcel               []ParcelInfo           `json:",omitempty,omitzero"`
	BiweeklyProductMail                 []ParcelInfo           `json:",omitempty,omitzero"`
	WeeklyProductParcel                 []ParcelInfo           `json:",omitempty,omitzero"`
	WeeklyProductMail                   []ParcelInfo           `json:",omitempty,omitzero"`
	EncryptedUID                        string                 `json:",omitempty,omitzero"`
	AccountRestrictionsDB               AccountRestrictionsDB  `json:",omitempty,omitzero"`
	IssueAlertInfos                     []IssueAlertInfoDB     `json:",omitempty,omitzero"`
	DailyRecordDBs                      []DailyRecordDB        `json:",omitempty,omitzero"`
	OptionDB                            OptionDB               `json:",omitempty,omitzero"`
	IsArenaAnonymous                    bool                   `json:",omitempty,omitzero"`
}
