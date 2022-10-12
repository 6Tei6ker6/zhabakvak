package toads

import (
	"fmt"
)

type Toad struct {
	Name	string
	Id 		string
	HP 		int
	Dmg 	int 
	Speed 	int
	Defence int
	CrtDmg	float32
	CrtChan	
}

func (t *Toad) IsAlive() (bool, error) {
	if t.HP <= 0 {
		t.HP = 0 
		return false, nil
	}
	return true, nil
}


