package collection

var systemType = map[int]string{
	1: "queue",
	2: "direct",
	3: "backup_system",
}

// SystemType ...
func SystemType(systemTypeID int) string {
	return systemType[systemTypeID]
}
