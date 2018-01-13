package npmdl

import "encoding/json"

// Point client for point api.
type Point struct {
	*Client
}

// NewPoint client.
func NewPoint() *Point {
	return &Point{
		Client: &Client{},
	}
}

// PointResponse returns information about point download counts
type PointResponse struct {
	Downloads int    `json:"downloads"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Package   string `json:"package"`
}

// GetCounts gets the point counts for the pacakge for the given period
// pkg is the package name
// optionally pass in a period (default: "last-day")
func (c *Point) GetCounts(pkg string, period ...string) (out *PointResponse, err error) {
	p := "last-day"
	if len(period) > 0 {
		p = period[0]
	}

	body, err := c.call("/point/" + pkg + "/" + p)
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
