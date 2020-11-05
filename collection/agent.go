package collection

var agentType = map[int]string{
	1: "Call Taker",
	2: "Non Emergency",
	3: "Coordinator",
	4: "Dispatcher",
	5: "Supervisor",
	7: "Call Taker and Non Emergency Swarm",
	8: "Coordinator and Dispatcher Swarm",
}

// AgentType ....
func AgentType(agentTypeID int) string {
	return agentType[agentTypeID]
}
