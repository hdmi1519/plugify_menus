package constants

const (
	IN_ATTACK    uint64 = 1 << 0  // LMB
	IN_JUMP      uint64 = 1 << 1  // Space
	IN_DUCK      uint64 = 1 << 2  // Ctrl
	IN_FORWARD   uint64 = 1 << 3  // W
	IN_BACK      uint64 = 1 << 4  // S
	IN_USE       uint64 = 1 << 5  // E
	IN_MOVELEFT  uint64 = 1 << 9  // A
	IN_MOVERIGHT uint64 = 1 << 10 // D
	IN_ATTACK2   uint64 = 1 << 11 // RMB
	IN_RELOAD    uint64 = 1 << 13 // R
	IN_SCORE     uint64 = 1 << 33 // Tab
)

const MenuButtons = IN_FORWARD | IN_BACK | IN_USE | IN_MOVELEFT | IN_MOVERIGHT | IN_RELOAD
