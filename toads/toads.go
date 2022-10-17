package toads

type Toad struct {
	Name      string
	Id        string
	SpellList []string
	HP        int     // 60~
	Dmg       int     // 40~
	Speed     int     // 35~
	Defence   int     // 60~
	CrtDmg    float32 // 2.0~
	CrtChan   float32 // 0.05~
}

func (t *Toad) IsAlive() (bool, error) {
	if t.HP <= 0 {
		t.HP = 0
		return false, nil
	}
	return true, nil
}
