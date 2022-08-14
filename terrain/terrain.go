package terrain

type Terrain struct {
	PeakMsg     string
	CantMoveMsg string
	DeathMsg    string

	// Stepping on the terrain changes the game state
	GameState int
}

func (t Terrain) IsFree() bool {
	return t.CantMoveMsg == ""
}

func (t Terrain) IsDeadly() bool {
	return t.DeathMsg != ""
}
