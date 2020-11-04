package collection

var sourceType = map[int]string{
	1: "Cellular",
	2: "Data",
	3: "TTRS",
	4: "191",
	5: "Operator",
}

// SourceType ...
func SourceType(sourceTypeID int) string {
	return sourceType[sourceTypeID]
}
