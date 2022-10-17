package protect

import (
	"zhabakvak/buffs"
)

type DefaultProtect struct {
	Defence int
}

func NewDefaultProtect(baseDefence int) *DefaultProtect {
	return &DefaultProtect{Defence: baseDefence}
}

func (d *DefaultProtect) Calculate(dmg int, buff []buffs.Buff) int {
	var fbuff int // final buff

	for _, x := range buff {
		if x.Type == buffs.DefenceBuff {
			fbuff = fbuff + x.Impact
		}
	}

	fdmg := dmg - d.Defence + fbuff

	if fdmg <= 5 {
		fdmg = 5
		return fdmg
	}

	return fdmg
}
