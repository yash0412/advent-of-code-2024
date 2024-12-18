package models

type Coords struct {
	X int
	Y int
}

func (c Coords) IsAtPOS(pos Coords) bool {
	return c.X == pos.X && c.Y == pos.Y
}

func (c Coords) IsWithinBounds(width, height int) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < width && c.Y < height
}

func (c Coords) MovePos(dx, dy int) Coords {
	return Coords{X: c.X + dx, Y: c.Y + dy}
}
