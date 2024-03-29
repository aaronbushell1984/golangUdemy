package decoupling

import "testing"

func TestPerimeter(t *testing.T) {
	r1 := Rectangle{
		height: 10.0,
		width:  10.0,
	}
	got := Perimeter(r1)
	want := 40.0
	if got != want {
		t.Errorf("got: %.2f want: %.2f", got, want)
	}
}

func TestArea_helper(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got: %g want: %g", got, want)
		}
	}
	t.Run("test rectangle area", func(t *testing.T) {
		r1 := Rectangle{
			height: 10.0,
			width:  10.0,
		}
		checkArea(t, r1, 100.0)
	})
	t.Run("test circle area", func(t *testing.T) {
		c1 := Circle{
			radius: 10.0,
		}
		checkArea(t, c1, 314.1592653589793)
	})

}

func TestArea_table(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{10, 10}, 50.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v, got: %g want: %g", tt.shape, got, tt.want)
			}
		})
	}

}
