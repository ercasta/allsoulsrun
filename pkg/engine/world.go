package engine

type World struct {
	Characters []Character
}

func (w *World) AddCharacter(c Character) {
	w.Characters = append(w.Characters, c)
}
