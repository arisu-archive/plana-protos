package protos

type AcademyDB struct {
	LastUpdate               MxTime                          `json:",omitempty,omitzero"`
	ZoneVisitCharacterDBs    map[int64][]VisitingCharacterDB `json:",omitempty,omitzero"`
	ZoneScheduleGroupRecords map[int64][]int64               `json:",omitempty,omitzero"`
}
