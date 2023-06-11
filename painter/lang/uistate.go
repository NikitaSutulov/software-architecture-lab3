package lang

import (
	"github.com/NikitaSutulov/software-architecture-lab3/painter"
)

type Uistate struct {
	backgroundColor     painter.Operation
	backgroundRectangle *painter.BackgroundRectangle
	figuresArray        []*painter.CrossFigure
}
