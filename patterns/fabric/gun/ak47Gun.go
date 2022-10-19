package gun


// Ak47Gun ...
type Ak47Gun struct {
	Gun
}

func newAk47() IGun {
	return &Ak47Gun{
		Gun: Gun{
			name: "Ak47",
			power: 4,
		},
	}
}