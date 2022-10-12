package toads

import (
	"fmt"
)

func (t *Toad) LowKick(target *Toad) (string, error) {
	dmg := t.Dmg + 30
	postdmg, _ := target.DoDefence(dmg) // добавить обработку ошибки
	target.HP = target.HP - postdmg
	lose, err := target.HealthCheck() // добавить обработку ошибки
	if lose != "" {
		return lose, err
	}
	return fmt.Sprintf("Урон %v, у жабки %v осталось %v хп", postdmg, target.Name, target.HP), nil
}

func (t *Toad) SelfRegen() (string, error) {
	t.HP = t.HP + 10
	return fmt.Sprintf("хп жабки %v восстановлено до %v", t.Name, t.HP), nil
}