package collection

// Agent ...
type Agent struct {
	ID   int
	Type string
}

// AgentCollection ...
type AgentCollection struct {
	CallTaker                     Agent
	NonEmergency                  Agent
	Coordinator                   Agent
	Dispatcher                    Agent
	Supervisor                    Agent
	CallTakerAndNonEmergencySwarm Agent
	CoordinatorAndDispatcherSwarm Agent
}

// AgentType ....
func AgentType() AgentCollection {
	return AgentCollection{
		CallTaker:                     Agent{ID: 1, Type: "Call Taker"},
		NonEmergency:                  Agent{ID: 2, Type: "Non Emergency"},
		Coordinator:                   Agent{ID: 3, Type: "Coordinator"},
		Dispatcher:                    Agent{ID: 4, Type: "Dispatcher"},
		Supervisor:                    Agent{ID: 5, Type: "Supervisor"},
		CallTakerAndNonEmergencySwarm: Agent{ID: 6, Type: "Call Taker and Non Emergency Swarm"},
		CoordinatorAndDispatcherSwarm: Agent{ID: 7, Type: "Coordinator and Dispatcher Swarm"},
	}
}
