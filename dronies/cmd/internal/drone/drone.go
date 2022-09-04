package drone

import (
	"github.com/sony/sonyflake"
	"math/rand"
	"time"
)

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

var sf = sonyflake.NewSonyflake(sonyflake.Settings{
	StartTime: time.Now(),
	MachineID: func() (uint16, error) {
		return 1, nil
	},
})

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
