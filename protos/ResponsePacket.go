package protos

type ResponsePacket struct {
	BasePacket
	ServerTimeTicks                 int64                              `json:",omitempty,omitzero"`
	ServerNotification              ServerNotificationFlag             `json:",omitempty,omitzero"`
	MissionProgressDBs              []MissionProgressDB                `json:",omitempty,omitzero"`
	EventMissionProgressDBDict      map[int64][]MissionProgressDB      `json:",omitempty,omitzero"`
	BattlePassMissionProgressDBDict map[int64][]MissionProgressDB      `json:",omitempty,omitzero"`
	StaticOpenConditions            map[string]OpenConditionLockReason `json:",omitempty,omitzero"`
}
