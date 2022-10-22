package color

import (
	"fmt"
	"testing"
)

func TestColorEquality(t *testing.T) {
	var tests = []struct {
		c1, c2 Color
		want   bool
	}{
		{Color{.1, .2, .3}, Color{.1, .2, .3}, true},
		{Color{0, .2, .3}, Color{.1, .2, .3}, false},
		{Color{.1, 0, .3}, Color{.1, .2, .3}, false},
		{Color{.1, .2, 0}, Color{.1, .2, .3}, false},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.c1, tt.c2)
		t.Run(name, func(t *testing.T) {
			ans := Equal(tt.c1, tt.c2)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestAddColors(t *testing.T) {
	c1 := Color{0.1, 0.2, 0.3}
	c2 := Color{0.4, 0.5, 0.6}
	got := c1.Add(c2)
	want := Color{0.5, 0.7, 0.9}
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestSubtractColors(t *testing.T) {
	c1 := Color{0.4, 0.5, 0.6}
	c2 := Color{0.1, 0.2, 0.3}
	got := c1.Subtract(c2)
	want := Color{0.3, 0.3, 0.3}
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMultiplyColorByScalar(t *testing.T) {
	c := Color{0.1, 0.2, 0.3}
	scalar := 2.0
	got := c.Multiply(scalar)
	want := Color{0.2, 0.4, 0.6}
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestPiecewiseMultiplyColors(t *testing.T) {
	c1 := Color{0.4, 0.5, 0.6}
	c2 := Color{0.1, 0.2, 0.3}
	got := PiecewiseMultiply(c1, c2)
	want := Color{0.04, 0.1, 0.18}
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
