package create_order

type EncounterInstanceId struct {
	Id          int64
	EncounterId int64
}

type CompleteEncounterCommandType int8

const (
	UpdateEncounter CompleteEncounterCommandType = iota
	RollbackEncounter
	UpdateFollower
	RollbackFollower
	UnknownCommand
)

type UpdateEncounterCommand struct {
	EncounterInstanceId EncounterInstanceId
	Type                CompleteEncounterCommandType
}

type CompleteEncounterReplyType int8

const (
	EncounterUpdated CompleteEncounterReplyType = iota
	EncounetrNotUpdated
	EncounterRolledBack
	FollowerUpdated
	FollowerNotUpdated
	FollowerRolledBack
	UnknownReply
)

type CompleteEncounterReply struct {
	EncounterInstanceId EncounterInstanceId
	Type                CompleteEncounterReplyType
}
