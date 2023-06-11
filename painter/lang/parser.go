package lang

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"strconv"
	"strings"

	"github.com/NikitaSutulov/software-architecture-lab3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	uistate         Uistate
	moveOperations  []painter.Operation
	updateOperation painter.Operation
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	if p.uistate.backgroundColor == nil {
		p.uistate.backgroundColor = painter.OperationFunc(painter.Reset)
	}
	if p.updateOperation != nil {
		p.updateOperation = nil
	}

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cmdl := scanner.Text()

		err := p.parse(cmdl)
		if err != nil {
			return nil, err
		}
	}

	var res []painter.Operation

	if p.uistate.backgroundColor != nil {
		res = append(res, p.uistate.backgroundColor)
	}
	if p.uistate.backgroundRectangle != nil {
		res = append(res, p.uistate.backgroundRectangle)
	}
	if len(p.moveOperations) != 0 {
		res = append(res, p.moveOperations...)
		p.moveOperations = nil
	}
	if len(p.uistate.figuresArray) != 0 {
		for _, figure := range p.uistate.figuresArray {
			res = append(res, figure)
		}
	}
	if p.updateOperation != nil {
		res = append(res, p.updateOperation)
	}

	return res, nil
}

func (p *Parser) reset() {
	p.uistate.backgroundColor = nil
	p.uistate.backgroundRectangle = nil
	p.uistate.figuresArray = nil
	p.moveOperations = nil
	p.updateOperation = nil
}

func (p *Parser) parse(cmdl string) error {
	words := strings.Split(cmdl, " ")
	command := words[0]

	switch command {
	case "white":
		if len(words) != 1 {
			return fmt.Errorf("wrong number of arguments for white command")
		}
		p.uistate.backgroundColor = painter.OperationFunc(painter.WhiteFill)
	case "green":
		if len(words) != 1 {
			return fmt.Errorf("wrong number of arguments for green command")
		}
		p.uistate.backgroundColor = painter.OperationFunc(painter.GreenFill)
	case "bgrect":
		parameters, err := checkForErrorsInParameters(words, 5)
		if err != nil {
			return err
		}
		p.uistate.backgroundRectangle = &painter.BackgroundRectangle{
			FirstPoint:  image.Point{X: parameters[0], Y: parameters[1]},
			SecondPoint: image.Point{X: parameters[2], Y: parameters[3]},
		}
	case "figure":
		parameters, err := checkForErrorsInParameters(words, 3)
		if err != nil {
			return err
		}

		figure := painter.CrossFigure{
			CentralPoint: image.Point{X: parameters[0], Y: parameters[1]},
		}
		p.uistate.figuresArray = append(p.uistate.figuresArray, &figure)
	case "move":
		parameters, err := checkForErrorsInParameters(words, 3)
		if err != nil {
			return err
		}
		moveOp := painter.MoveOperation{X: parameters[0], Y: parameters[1], FiguresArray: p.uistate.figuresArray}
		p.moveOperations = append(p.moveOperations, &moveOp)
	case "reset":
		if len(words) != 1 {
			return fmt.Errorf("wrong number of arguments for reset command")
		}
		p.reset()
		p.uistate.backgroundColor = painter.OperationFunc(painter.Reset)
	case "update":
		if len(words) != 1 {
			return fmt.Errorf("wrong number of arguments for update command")
		}
		p.updateOperation = painter.UpdateOp
	default:
		return fmt.Errorf("invalid command %v", words[0])
	}
	return nil
}

func checkForErrorsInParameters(words []string, expected int) ([]int, error) {
	if len(words) != expected {
		return nil, fmt.Errorf("wrong number of arguments for '%v' command", words[0])
	}
	var command = words[0]
	var params []int
	for _, param := range words[1:] {
		p, err := parseInt(param)
		if err != nil {
			return nil, fmt.Errorf("invalid parameter for '%s' command: '%s' is not a number", command, param)
		}
		params = append(params, p)
	}
	return params, nil
}

func parseInt(s string) (int, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot parse float: %s", s)
	}
	return int(f * 800), nil
}
