package collection

// System ...
type System struct {
	ID   int
	Type string
}

// SystemCollection ...
type SystemCollection struct {
	Queue        System
	Direct       System
	BackupSystem System
}

// SystemType ...
func SystemType() SystemCollection {
	return SystemCollection{
		Queue:        System{ID: 1, Type: "queue"},
		Direct:       System{ID: 2, Type: "direct"},
		BackupSystem: System{ID: 3, Type: "backup_system"},
	}
}
