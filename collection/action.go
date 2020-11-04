package collection

// Action ...
type Action struct {
	ID   int
	Type string
}

// ActionCollection ...
type ActionCollection struct {
	Process                 Action
	EnterQueue              Action
	Ringing                 Action
	Answer                  Action
	Hangup                  Action
	Transfer                Action
	RingCanceled            Action
	QueueFullAnother        Action
	CreateConference        Action
	JoinConference          Action
	LeaveConference         Action
	SwarmNormal             Action
	SwarmEmer               Action
	TransferQueue           Action
	TransferJoinCalling     Action
	Pull                    Action
	Spy                     Action
	NoAnswer                Action
	Abandon                 Action
	QueueFullAbandon        Action
	QueueFullAnotherAbandon Action
	DNDOff                  Action
	DNDShort                Action
	DNDLong                 Action
	PauseAgent              Action
	Offline                 Action
}

// ActionType ...
func ActionType() ActionCollection {
	return ActionCollection{
		Process:                 Action{ID: 101, Type: "PROCESS"},
		EnterQueue:              Action{ID: 102, Type: "ENTERQUEUE"},
		Ringing:                 Action{ID: 201, Type: "RINGING"},
		Answer:                  Action{ID: 202, Type: "ANSWER"},
		Hangup:                  Action{ID: 203, Type: "HANGUP"},
		Transfer:                Action{ID: 204, Type: "TRANSFER"},
		RingCanceled:            Action{ID: 205, Type: "RINGCANCELED"},
		QueueFullAnother:        Action{ID: 206, Type: "QUEUE_FULL_ANOTHER"},
		CreateConference:        Action{ID: 207, Type: "CREATE_CONFERENCE"},
		JoinConference:          Action{ID: 208, Type: "JOIN_CONFERENCE"},
		LeaveConference:         Action{ID: 209, Type: "LEAVE_CONFERENCE"},
		SwarmNormal:             Action{ID: 211, Type: "SWARM_NORMAL"},
		SwarmEmer:               Action{ID: 212, Type: "SWARM_EMER"},
		TransferQueue:           Action{ID: 213, Type: "TRANSFER_QUEUE"},
		TransferJoinCalling:     Action{ID: 214, Type: "TRANSFER_JOIN_CALLING"},
		Pull:                    Action{ID: 215, Type: "PULL"},
		Spy:                     Action{ID: 216, Type: "SPY"},
		NoAnswer:                Action{ID: 301, Type: "NO_ANSWER"},
		Abandon:                 Action{ID: 302, Type: "ABANDON"},
		QueueFullAbandon:        Action{ID: 303, Type: "QueueFullAbandon"},
		QueueFullAnotherAbandon: Action{ID: 304, Type: "QUEUE_FULL_ANOTHER_ABANDON"},
		DNDOff:                  Action{ID: 401, Type: "DND_OFF"},
		DNDShort:                Action{ID: 402, Type: "DND_SHORT"},
		DNDLong:                 Action{ID: 403, Type: "DND_LONG"},
		PauseAgent:              Action{ID: 404, Type: "PAUSE_AGENT"},
		Offline:                 Action{ID: 405, Type: "OFFLINE"},
	}
}
