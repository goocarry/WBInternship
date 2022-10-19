package gun

// MusketGun ...
type MusketGun struct {
	Gun
}

func newMusket() IGun {
	return &MusketGun{
		Gun: Gun{
			name: "Musket",
			power: 1,
		},
	}
}