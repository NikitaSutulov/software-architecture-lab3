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
	uistate Uistate
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	p.uistate.ResetOperations()

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cmdl := scanner.Text()

		err := p.parse(cmdl)
		if err != nil {
			return nil, err
		}
	}

	res := p.uistate.GetOperations()

	return res, nil
}

func (p *Parser) parse(cmdl string) error {
	words := strings.Split(cmdl, " ")
	command := words[0]

	switch command {
	case "white":
		if len(words) != 1 {
			return fmt.Errorf("wrong number of arguments for white command")
		}
		p.uistate.WhiteBackground()
	case "green":
		if len(words) != 1 {
			return fmt.Errorf("wrong number of arguments for green command")
		}
		p.uistate.GreenBackground()
	case "bgrect":
		parameters, err := checkForErrorsInParameters(words, 5)
		if err != nil {
			return err
		}
		p.uistate.BackgroundRectangle(image.Point{X: parameters[0], Y: parameters[1]}, image.Point{X: parameters[2], Y: parameters[3]})
	case "figure":
		parameters, err := checkForErrorsInParameters(words, 3)
		if err != nil {
			return err
		}

		p.uistate.AddFigure(image.Point{X: parameters[0], Y: parameters[1]})
	case "move":
		parameters, err := checkForErrorsInParameters(words, 3)
		if err != nil {
			return err
		}
		p.uistate.AddMoveOperation(parameters[0], parameters[1])
	case "reset":
		if len(words) != 1 {
			return fmt.Errorf("wrong number of arguments for reset command")
		}
		p.uistate.ResetStateAndBackground()
	case "update":
		if len(words) != 1 {
			return fmt.Errorf("wrong number of arguments for update command")
		}
		p.uistate.SetUpdateOperation()
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
