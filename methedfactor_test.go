package useage

import (
	"testing"
)

func TestOne(t *testing.T) {
	area := new(Text)
	area.SetX(10)
	if area.GetX() != 10 {
		t.Error("test methodfactor:One failed")
	} else {
		t.Log("test methodfactor:One pass")
	}
}

func TestTwo(t *testing.T) {
	area := *new(Text)
	area.SetY(20)
	if area.GetY() != 20 {
		t.Error("test methodfactor:Two failed")
	} else {
		t.Log("test methodfactor:Two pass")
	}
}
