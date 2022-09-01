package terrain

type Terrain struct {
	PeakMsg     string
	CantMoveMsg string
	DeathMsg    string

	// Stepping on the terrain changes the game state
	GameState int
	// String that can be used by actions to tell terrain apart
	ActionStr string
}

func (t Terrain) IsFree() bool {
	return t.CantMoveMsg == ""
}

func (t Terrain) IsDeadly() bool {
	return t.DeathMsg != ""
}
