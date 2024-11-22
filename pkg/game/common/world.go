package common

import (
	e "github.com/ercasta/allsoulsrun/pkg/engine"
)

const WorldComponentType e.ComponentType = "World"

type World struct {
	Characters []e.EntityID
}

func (w *World) AddCharacter(c e.EntityID) {
	w.Characters = append(w.Characters, c)
}

func (w World) GetComponentType() e.ComponentType { return WorldComponentType }
