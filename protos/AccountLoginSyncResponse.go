package protos

type AccountLoginSyncResponse struct {
	ResponsePacket
	Responses                                ResponsePacket
	CafeGetInfoResponse                      CafeGetInfoResponse
	AccountCurrencySyncResponse              AccountCurrencySyncResponse
	CharacterListResponse                    CharacterListResponse
	EquipmentItemListResponse                EquipmentItemListResponse
	CharacterGearListResponse                CharacterGearListResponse
	ItemListResponse                         ItemListResponse
	EchelonListResponse                      EchelonListResponse
	MemoryLobbyListResponse                  MemoryLobbyListResponse
	CampaignListResponse                     CampaignListResponse
	ArenaLoginResponse                       ArenaLoginResponse
	RaidLoginResponse                        RaidLoginResponse
	EliminateRaidLoginResponse               EliminateRaidLoginResponse
	CraftInfoListResponse                    CraftInfoListResponse
	ClanLoginResponse                        ClanLoginResponse
	MomotalkOutlineResponse                  MomoTalkOutLineResponse
	ScenarioListResponse                     ScenarioListResponse
	ShopGachaRecruitListResponse             ShopGachaRecruitListResponse
	TimeAttackDungeonLoginResponse           TimeAttackDungeonLoginResponse
	BillingPurchaseListByYostarResponse      BillingPurchaseListByYostarResponse
	EventContentPermanentListResponse        EventContentPermanentListResponse
	AttachmentGetResponse                    AttachmentGetResponse
	AttachmentEmblemListResponse             AttachmentEmblemListResponse
	ContentSweepMultiSweepPresetListResponse ContentSweepMultiSweepPresetListResponse
	StickerListResponse                      StickerLoginResponse
	MultiFloorRaidSyncResponse               MultiFloorRaidSyncResponse
	MultiFloorRaidLoginResponse              MultiFloorRaidLoginResponse
	FriendCount                              int64 `json:",omitempty,omitzero"`
	FriendCode                               string
	PickupFirstGetHistoryDBs                 []PickupFirstGetHistoryDB
	AccountLevelRewardIds                    []int64
}
