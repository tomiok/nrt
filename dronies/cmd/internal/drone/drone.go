package drone

import (
	"github.com/sony/sonyflake"
	"math/rand"
	"time"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{
	StartTime: time.Now(),
	MachineID: func() (uint16, error) {
		return 1, nil
	},
})

type Drone struct {
	ID  uint64
	Lat float64
	Lon float64

	Scans Scan
}

type Scan struct {
	ScanTime time.Time
	Squad    Squad
}

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

func DronesGeneration(c int) []Drone {
	drones := make([]Drone, 0, c)
	for i := 0; i < c; i++ {
		drone, err := createDrone()
		if err != nil {
			i--
			continue
		}
		drones = append(drones, drone)
	}
	return drones
}

func createDrone() (Drone, error) {
	id, err := sf.NextID()

	if err != nil {
		return Drone{}, err
	}

	return Drone{
		ID:  id,
		Lat: randFloats(-90.00, 90.00),
		Lon: randFloats(-180.00, 180.00),
	}, nil

}

func randFloats(min, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}

var random = rand.New(rand.NewSource(time.Now().Unix()))

func (d *Drone) ScanEnemy() {
	var squad Squad
	enemies := make([]Enemy, 0, 100)

	var randEnemies = random.Intn(110)
	for i := 0; i < randEnemies; i++ {
		enemies = append(enemies, Enemy{
			ID:     generateID(),
			Energy: random.Intn(100),
			Skill:  randSkill(random),
		})
	}
	squad.Enemies = enemies
	d.Scans = Scan{
		ScanTime: time.Now(),
		Squad:    squad,
	}
}

func (d *Drone) Read() Message {
	droneID := d.ID
	readAt := d.Scans.ScanTime.Format(time.ANSIC)
	pos := Position{
		Lon: d.Lon,
		Lat: d.Lat,
	}
	enemies := d.Scans.Squad.Enemies

	return Message{
		DroneID:  droneID,
		Position: pos,
		ScanAt:   readAt,
		Enemies:  enemies,
	}
}
