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
	IsActive() bool
	Toggle()
	Clear()
}

type baseElement struct {
	id       ElementID
	bounds   *geometry.Rect
	isActive bool
}

func (b *baseElement) ID() ElementID            { return b.id }
func (b *baseElement) Bounds() *geometry.Rect   { return b.bounds }
func (b *baseElement) Contains(cx, cy int) bool { return b.Bounds().Contains(cx, cy) }
func (b *baseElement) IsActive() bool           { return b.isActive }
func (b *baseElement) Toggle()                  { b.isActive = !b.isActive }
func (b *baseElement) Clear()                   { b.isActive = true }
