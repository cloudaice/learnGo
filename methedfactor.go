package useage

type Text struct {
    x int
    y int
}

func (t *Text) SetX(v int) {
    t.x = v
}

func (t *Text) GetX() int{
    return t.x
}

func (t *Text) SetY(v int) {
    t.y = v
}

func (t *Text) GetY() int {
    return t.y
}

