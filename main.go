package main

import (
	"zhabakvak/combat"
	"zhabakvak/toads"
)

var(
Zhab = toads.Toad{Name: "Huesos", HP: 60, Dmg: 40, Speed:35, Defence: 60}
Kvab = &toads.Toad{Name: "Muertos", HP: 60, Dmg: 40, Speed:35, Defence: 60}
)

func main() {
	combat.Drasting(Zhab, Kvab)
}