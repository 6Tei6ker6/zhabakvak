package buffs

type Buff struct {
	Type   string
	Impact int
}

const (
	HpBuff      = "HP"
	DmgBuff     = "DMG"
	SpeedBuff   = "SPEED"
	DefenceBuff = "DEFENCE"
	CrtDmgBuff  = "CRTDMG"
	CtrChanBuff = "CRTCHAN"
)
