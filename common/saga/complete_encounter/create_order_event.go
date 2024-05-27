package create_order

type EncounterInstanceId struct {
	Id          int64
	EncounterId int64
}

type CompleteEncounterCommandType int8

const (
	UpdateFollower CompleteEncounterCommandType = iota
	RollbackFollower
	RollbackEncounter
	UnknownCommand
	CompleteEncounter
)

type UpdateEncounterCommand struct {
	EncounterInstanceId EncounterInstanceId
	Type                CompleteEncounterCommandType
}

type CompleteEncounterReplyType int8

const (
	FollowerUpdated CompleteEncounterReplyType = iota
	FollowerNotUpdated
	FollowerRolledBack
	UnknownReply
	EncounterUpdated
	EncounterCompleted
)

type CompleteEncounterReply struct {
	EncounterInstanceId EncounterInstanceId
	Type                CompleteEncounterReplyType
}
