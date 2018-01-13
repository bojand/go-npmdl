package npmdl

// PointCounts contains point download counts data
type PointCounts struct {
	Downloads int    `json:"downloads"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Package   string `json:"package"`
}

// DayCounts contains download count data for a specific day
type DayCounts struct {
	Downloads int    `json:"downloads"`
	Day       string `json:"day"`
}

// RangeCounts contains download count data for a specific range
type RangeCounts struct {
	Start     string       `json:"start"`
	End       string       `json:"end"`
	Package   string       `json:"package"`
	Downloads []*DayCounts `json:"downloads"`
}
