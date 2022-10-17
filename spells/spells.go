package spells 

import (
	"zhabakvak/buffs"
)

type Spell interface {
	Calculate(dmg int, buff []buffs.Buff) (odmg, hpgain int, obuff []buffs.Buff)
}

var (
	Spells map[string]Spell
)
