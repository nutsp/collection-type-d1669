package collection

type Source struct {
	ID   int
	Type string
}

// SourceCollection ...
type SourceCollection struct {
	Cellular Source
	Data     Source
	TTRS     Source
	N191     Source
	Operator Source
}

// SourceType ...
func SourceType() SourceCollection {
	return SourceCollection{
		Cellular: Source{ID: 1, Type: "Cellular"},
		Data:     Source{ID: 2, Type: "Data"},
		TTRS:     Source{ID: 3, Type: "TTRS"},
		N191:     Source{ID: 4, Type: "191"},
		Operator: Source{ID: 5, Type: "Operator"},
	}
}
