package spells

import (
	"zhabakvak/buffs"
)

type SelfCare struct {
	health 	int //20
}

func NewSelfCare(baseHealth int) *SelfCare {
	return &SelfCare{health: baseHealth}
}

func (s *SelfCare) Calculate (dmg int, buff []buffs.Buff) (odmg, hpgain int, obuff []buffs.Buff) {
	return 	
}