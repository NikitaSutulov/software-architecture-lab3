package lang

import (
	"github.com/NikitaSutulov/software-architecture-lab3/painter"
	"image"
)

type Uistate struct {
	backgroundColor     painter.Operation
	backgroundRectangle *painter.BackgroundRectangle
	figuresArray        []*painter.CrossFigure
	moveOperations      []painter.Operation
	updateOperation     painter.Operation
}

func (u *Uistate) Reset() {
	u.backgroundColor = nil
	u.backgroundRectangle = nil
	u.figuresArray = nil
	u.moveOperations = nil
	u.updateOperation = nil
}

func (u *Uistate) GetOperations() []painter.Operation {
	var ops []painter.Operation

	if u.backgroundColor != nil {
		ops = append(ops, u.backgroundColor)
	}
	if u.backgroundRectangle != nil {
		ops = append(ops, u.backgroundRectangle)
	}
	if len(u.moveOperations) != 0 {
		ops = append(ops, u.moveOperations...)
		u.moveOperations = nil
	}
	if len(u.figuresArray) != 0 {
		for _, figure := range u.figuresArray {
			ops = append(ops, figure)
		}
	}
	if u.updateOperation != nil {
		ops = append(ops, u.updateOperation)
	}

	return ops
}

func (u *Uistate) ResetOperations() {
	if u.backgroundColor == nil {
		u.backgroundColor = painter.OperationFunc(painter.Reset)
	}
	if u.updateOperation != nil {
		u.updateOperation = nil
	}
}

func (u *Uistate) GreenBackground() {
	u.backgroundColor = painter.OperationFunc(painter.GreenFill)
}

func (u *Uistate) WhiteBackground() {
	u.backgroundColor = painter.OperationFunc(painter.GreenFill)
}

func (u *Uistate) BackgroundRectangle(firstPoint image.Point, secondPoint image.Point) {
	u.backgroundRectangle = &painter.BackgroundRectangle{
		FirstPoint:  firstPoint,
		SecondPoint: secondPoint,
	}
}

func (u *Uistate) AddFigure(centralPoint image.Point) {
	figure := painter.CrossFigure{
		CentralPoint: centralPoint,
	}
	u.figuresArray = append(u.figuresArray, &figure)
}

func (u *Uistate) AddMoveOperation(x int, y int) {
	moveOp := painter.MoveOperation{X: x, Y: y, FiguresArray: u.figuresArray}
	u.moveOperations = append(u.moveOperations, &moveOp)
}

func (u *Uistate) ResetStateAndBackground() {
	u.Reset()
	u.backgroundColor = painter.OperationFunc(painter.Reset)
}

func (u *Uistate) SetUpdateOperation() {
	u.updateOperation = painter.UpdateOp
}
