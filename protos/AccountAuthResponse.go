package protos

type AccountAuthResponse struct {
	ResponsePacket
	CurrentVersion                      int64 `json:",omitempty,omitzero"`
	MinimumVersion                      int64 `json:",omitempty,omitzero"`
	IsDevelopment                       bool  `json:",omitempty,omitzero"`
	BattleValidation                    bool  `json:",omitempty,omitzero"`
	UpdateRequired                      bool  `json:",omitempty,omitzero"`
	TTSCdnUri                           string
	AccountDB                           AccountDB
	AttendanceBookRewards               []AttendanceBookReward
	AttendanceHistoryDBs                []AttendanceHistoryDB
	RepurchasableMonthlyProductCountDBs []PurchaseCountDB
	MonthlyProductParcel                []ParcelInfo
	MonthlyProductMail                  []ParcelInfo
	BiweeklyProductParcel               []ParcelInfo
	BiweeklyProductMail                 []ParcelInfo
	WeeklyProductParcel                 []ParcelInfo
	WeeklyProductMail                   []ParcelInfo
	EncryptedUID                        string
	AccountRestrictionsDB               AccountRestrictionsDB
	IssueAlertInfos                     []IssueAlertInfoDB
	DailyRecordDBs                      []DailyRecordDB
	OptionDB                            OptionDB
	IsArenaAnonymous                    bool `json:",omitempty,omitzero"`
	AccountLimitedFlashSaleDBs          []AccountLimitedFlashSaleDB
	NewlyAddedShopCashIds               []int64
}
