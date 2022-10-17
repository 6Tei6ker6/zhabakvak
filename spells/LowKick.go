package spells

import (
	"zhabakvak/buffs"
)

type LowKick struct {
	dmg int // 30
}

func NewLowKick(baseDmg int) *LowKick{
	return &LowKick{dmg: baseDmg}
}

func (s *LowKick) Calculate(dmg int, buff []buffs.Buff) (odmg, hpgain int, obuff []buffs.Buff) {
	var fbuff int // final_buff
	for _, x := range buff {
		if x.Type == buffs.DmgBuff {
			fbuff = fbuff + x.Impact
		}
	}
	odmg = s.dmg + dmg + fbuff 	
	return odmg, 0, nil //output_damage
}

func init() {
	Spells["LowKick"] = &LowKick{dmg: 30}
}

