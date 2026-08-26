package protos

type OpenConditionErrorPacket struct {
	ResponsePacket
	ErrorCode  WebAPIErrorCode         `json:",omitempty,omitzero"`
	LockReason OpenConditionLockReason `json:",omitempty,omitzero"`
}
