package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

func (b *box) checkIndex(i int) error {
	if i >= len(b.shapes) {
		return fmt.Errorf("Index doesn't exist")
	}

	if b.shapes[i] == nil {
		return fmt.Errorf("Empty item")
	}

	return nil
}

func (b *box) removeAt(i int) {
	if i < len(b.shapes) {
		b.shapes = append(b.shapes[:i], b.shapes[i+1:len(b.shapes)]...)
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}

	return fmt.Errorf("Box is full")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	err := b.checkIndex(i)

	if err != nil {
		return nil, err
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	err := b.checkIndex(i)

	if err != nil {
		return nil, err
	}

	shape := b.shapes[i]
	b.removeAt(i)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	err := b.checkIndex(i)

	if err != nil {
		return nil, err
	}

	old := b.shapes[i]
	b.shapes[i] = shape

	return old, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}

	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}

	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var found bool
	other := make([]Shape, 0)

	for _, shape := range b.shapes {
		switch shape.(type) {
		case *Circle:
			found = true
		default:
			other = append(other, shape)
		}
	}

	if !found {
		return fmt.Errorf("No circles found")
	}

	b.shapes = other

	return nil
}
