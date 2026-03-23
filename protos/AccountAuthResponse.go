package protos

type AccountAuthResponse struct {
	ResponsePacket
	CurrentVersion                      int64 `json:",omitempty,omitzero"`
	MinimumVersion                      int64 `json:",omitempty,omitzero"`
	IsDevelopment                       bool  `json:",omitempty,omitzero"`
	BattleValidation                    bool  `json:",omitempty,omitzero"`
	UpdateRequired                      bool  `json:",omitempty,omitzero"`
	TTSCdnUri                           string
	AccountDB                           *AccountDB `json:",omitempty,omitzero"`
	AttendanceBookRewards               []*AttendanceBookReward
	AttendanceHistoryDBs                []*AttendanceHistoryDB
	RepurchasableMonthlyProductCountDBs []*PurchaseCountDB
	MonthlyProductParcel                []*ParcelInfo
	MonthlyProductMail                  []*ParcelInfo
	BiweeklyProductParcel               []*ParcelInfo
	BiweeklyProductMail                 []*ParcelInfo
	WeeklyProductParcel                 []*ParcelInfo
	WeeklyProductMail                   []*ParcelInfo
	EncryptedUID                        string
	AccountRestrictionsDB               *AccountRestrictionsDB `json:",omitempty,omitzero"`
	IssueAlertInfos                     []*IssueAlertInfoDB
	DailyRecordDBs                      []*DailyRecordDB
	OptionDB                            *OptionDB `json:",omitempty,omitzero"`
	IsArenaAnonymous                    bool      `json:",omitempty,omitzero"`
	AccountLimitedFlashSaleDBs          []*AccountLimitedFlashSaleDB
	NewlyAddedShopCashIds               []int64
}
