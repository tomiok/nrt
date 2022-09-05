package monitor

type Position struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Message struct {
	DroneID  uint64   `json:"droneId"`
	Position Position `json:"position"`
	ScanAt   string   `json:"scanAt"`
	Enemies  []Enemy  `json:"enemies"`
}

type Enemy struct {
	ID     string `json:"id"`
	Energy int    `json:"energy"` // means a % of damage.
	Skill  string `json:"skill"`
}
