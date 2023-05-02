package painter

import (
	"image"
	"image/color"

	"golang.org/x/image/draw"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

// BackgroundRectangle структура чорного прямокутника
type BackgroundRectangle struct {
	FirstPoint  image.Point
	SecondPoint image.Point
}

// Do малює наш прямокутник

func (op *BackgroundRectangle) Do(t screen.Texture) bool {
	c := color.Black
	t.Fill(image.Rect(op.FirstPoint.X, op.FirstPoint.Y, op.SecondPoint.X, op.SecondPoint.Y), c, screen.Src)
	return false
}

type CrossFigure struct {
	CentralPoint image.Point
}

func (op *CrossFigure) Do(t screen.Texture) bool {
	c := color.RGBA{R: 255, G: 255, B: 0, A: 1}
	t.Fill(image.Rect(op.CentralPoint.X-200, op.CentralPoint.Y+80, op.CentralPoint.X+200, op.CentralPoint.Y-80), c, draw.Src)
	t.Fill(image.Rect(op.CentralPoint.X-80, op.CentralPoint.Y+200, op.CentralPoint.X+80, op.CentralPoint.Y-200), c, draw.Src)
	return false
}

type MoveOperation struct {
	X            int
	Y            int
	FiguresArray []*CrossFigure
}

func (op *MoveOperation) Do(t screen.Texture) bool {
	for i := range op.FiguresArray {
		op.FiguresArray[i].CentralPoint.X += op.X
		op.FiguresArray[i].CentralPoint.Y += op.Y
	}
	return false
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}
