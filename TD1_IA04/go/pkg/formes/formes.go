package formes

import (
	"math"
)

func Square(x int) int {
	return x * x
}

type Point2D struct {
	x float64
	y float64
}

func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{x, y}
}

func (p *Point2D) X() float64 {
	return p.x
}

func (p *Point2D) SetX(x float64) {
	p.x = x
}

func (p *Point2D) Y() float64 {
	return p.y
}

func (p *Point2D) SetY(y float64) {
	p.y = y
}

func (p *Point2D) Clone() *Point2D {
	return &Point2D{p.x, p.y}
}

func (p *Point2D) Module() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// ################################
// #          Rectangle          #
// ################################
type Rectangle struct {
	topLeft     *Point2D
	bottomRight *Point2D
}

func NewRectangle(topLeft, bottomRight *Point2D) *Rectangle {
	return &Rectangle{topLeft, bottomRight}
}

func (r *Rectangle) TopLeft() *Point2D {
	return r.topLeft
}

func (r *Rectangle) SetTopLeft(topLeft *Point2D) {
	r.topLeft = topLeft
}

func (r *Rectangle) BottomRight() *Point2D {
	return r.bottomRight
}

func (r *Rectangle) SetBottomRight(bottomRight *Point2D) {
	r.bottomRight = bottomRight
}

// ################################
// #           Sprite             #
// ################################

type Sprite struct {
	position *Point2D
	hitbox   *Rectangle
	zoom     float64
	filename string
}

func NewSprite(position *Point2D, hitbox *Rectangle, zoom float64, filename string) *Sprite {
	return &Sprite{position, hitbox, zoom, filename}
}

func (s *Sprite) Position() *Point2D {
	return s.position
}

func (s *Sprite) SetPosition(position *Point2D) {
	s.position = position
}

func (s *Sprite) Hitbox() *Rectangle {
	return s.hitbox
}

func (s *Sprite) SetHitbox(hitbox *Rectangle) {
	s.hitbox = hitbox
}

func (s *Sprite) Zoom() float64 {
	return s.zoom
}

func (s *Sprite) SetZoom(zoom float64) {
	s.zoom = zoom
}

func (s *Sprite) Filename() string {
	return s.filename
}

func (s *Sprite) SetFilename(filename string) {
	s.filename = filename
}

func (s *Sprite) Move(newPosition *Point2D) {
	s.position = newPosition
}

func (s *Sprite) Collision(other *Sprite) *Rectangle {
	leftX := max(s.hitbox.TopLeft().X(), other.hitbox.TopLeft().X())
	topY := max(s.hitbox.TopLeft().Y(), other.hitbox.TopLeft().Y())
	rightX := min(s.hitbox.BottomRight().X(), other.hitbox.BottomRight().X())
	bottomY := min(s.hitbox.BottomRight().Y(), other.hitbox.BottomRight().Y())

	if leftX < rightX && topY < bottomY {
		return NewRectangle(NewPoint2D(leftX, topY), NewPoint2D(rightX, bottomY))
	}

	return nil
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
