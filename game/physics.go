package game

type Location struct {
	X float64
	Y float64
}

func (l Location) Move(v Velocity) Location {
	return Location{
		X: l.X + v.X,
		Y: l.Y + v.Y,
	}
}

type Velocity struct {
	X float64
	Y float64
}

func (v Velocity) Avarage(other Velocity) Velocity {
	return Velocity{
		X: (v.X + other.X) / 2,
		Y: (v.Y + other.Y) / 2,
	}
}

func (v Velocity) Accelerate(p Power) Velocity {
	return Velocity{
		X: v.X + p.X,
		Y: v.Y + p.Y,
	}
}

type Power struct {
	X float64
	Y float64
}
