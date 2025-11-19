package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vhespanha/a_viagem/internal/geometry"
)

type ElementID uint8

type element interface {
	ID() ElementID
	Bounds() *geometry.Rect
	Contains(cx, cy int) bool
	Draw(screen *ebiten.Image)
	Active() bool
	Activate()
	Clear()
	Update(data any) error
	CustomClear()
}

type baseElement struct {
	id     ElementID
	bounds *geometry.Rect
	action CommandID
	data   any
	active bool
}

func (b *baseElement) ID() ElementID            { return b.id }
func (b *baseElement) Bounds() *geometry.Rect   { return b.bounds }
func (b *baseElement) Contains(cx, cy int) bool { return b.Bounds().Contains(cx, cy) }
func (b *baseElement) Active() bool             { return b.active }
func (b *baseElement) Activate()                { b.active = true }
func (b *baseElement) Clear()                   { b.active = false }
