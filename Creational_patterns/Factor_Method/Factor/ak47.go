package Factor

type ak47 struct {
	gun
}

func newAk47() IGun {
	return &ak47{
		gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
