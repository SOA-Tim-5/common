package create_order

type UpdateLevel struct {
	UserId string
	Level  string
}

type CompleteEncounterCommandType int8

const (
	UpdateFollower CompleteEncounterCommandType = iota
	RollbackFollower
	RollbackEncounter
	UnknownCommand
	CompleteEncounter
)

type UpdateLevelCommand struct {
	UpdateLevel UpdateLevel
	Type        CompleteEncounterCommandType
}

type CompleteEncounterReplyType int8

const (
	FollowerUpdated CompleteEncounterReplyType = iota
	FollowerNotUpdated
	FollowerRolledBack
	UnknownReply
	EncounterUpdated
	EncounterCompleted
	EncounterRolledBack
)

type CompleteEncounterReply struct {
	UpdateLevel UpdateLevel
	Type        CompleteEncounterReplyType
}
