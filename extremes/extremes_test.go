package extremes

import "testing"

func TestMinInt(t *testing.T) {

	a := 1
	b := 2
	exp := 1

	got := Min(a, b)
	if got != exp {
		t.Errorf("unexpected result, got: %d, want: %d", got, exp)
	}

}

func TestMaxInt(t *testing.T) {

	a := 1
	b := 2
	exp := 2

	got := Max(a, b)
	if got != exp {
		t.Errorf("unexpected result, got: %d, want: %d", got, exp)
	}

}

func TestMin64(t *testing.T) {

	var a int64 = 1
	var b int64 = 2
	var exp int64 = 1

	got := Min64(a, b)
	if got != exp {
		t.Errorf("unexpected result, got: %d, want: %d", got, exp)
	}

}

func TestMax64(t *testing.T) {

	var a int64 = 1
	var b int64 = 2
	var exp int64 = 2

	got := Max64(a, b)
	if got != exp {
		t.Errorf("unexpected result, got: %d, want: %d", got, exp)
	}

}
