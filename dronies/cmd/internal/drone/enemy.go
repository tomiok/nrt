package drone

import (
	"fmt"
	"github.com/rs/xid"
	"math/rand"
)

const (
	Easy      Skill = "easy"
	Medium    Skill = "medium"
	Hard      Skill = "hard"
	Nightmare Skill = "nightmare"
)

var skills = []Skill{Easy, Medium, Hard, Nightmare}

type Skill string

// Enemy is the one in the other side.
type Enemy struct {
	ID     string
	Energy int // means a % of damage.
	Skill  Skill
}

// Squad is simply a list of enemies.
type Squad struct {
	Enemies []Enemy
}

func randSkill(r *rand.Rand) Skill {
	idx := r.Intn(4)
	return skills[idx]
}

func generateID() string {
	return xid.New().String()
}

func (e Enemy) GetInfo() string {
	return fmt.Sprintf("ID:%s, energy:%d, skill:%s", e.ID, e.Energy, e.Skill)
}
