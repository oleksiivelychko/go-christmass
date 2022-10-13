package interfaces

import "testing"

func TestRectangleMeasure(t *testing.T) {
	r := Rectangle{width: 5, height: 5}

	area := measureArea(r)
	if area != 25 {
		t.Errorf("[func measureArea(g Geometry) float32] -> %f != 25", area)
	}

	perimeter := measurePerimeter(r)
	if perimeter != 20 {
		t.Errorf("[func measurePerimeter(g Geometry) float32] -> %f != 20", perimeter)
	}
}

func TestCircleMeasure(t *testing.T) {
	c := Circle{radius: 5}

	area := measureArea(c)
	if area != 78.539818 {
		t.Errorf("[func measureArea(g Geometry) float32] -> %f != 78.539818", area)
	}

	perimeter := measurePerimeter(c)
	if perimeter != 31.415928 {
		t.Errorf("[func measurePerimeter(g Geometry) float32] -> %f != 31.415928", perimeter)
	}
}
