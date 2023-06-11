package lang

import (
	"image"
	"strings"
	"testing"

	// Importing painter and lang packages.
	"github.com/NikitaSutulov/software-architecture-lab3/painter"

	// Package assert provides some helpful methods for testing, like for asserting equality or inequality.
	"github.com/stretchr/testify/assert"
	// Package require provides methods similar to assert but stop the test if the condition is not met, whereas assert will continue running.
	"github.com/stretchr/testify/require"
)

func Test_parse_struct(t *testing.T) {
	// Array of structs containing test cases.
	tests := []struct {
		name    string            // Name of the test case.
		command string            // The command to be parsed.
		op      painter.Operation // The expected painter.Operation struct after parsing.
	}{
		{
			name:    "background rectangle",
			command: "bgrect 0 0 0.125 0.125",
			op:      &painter.BackgroundRectangle{FirstPoint: image.Point{X: 0, Y: 0}, SecondPoint: image.Point{X: 100, Y: 100}},
		},
		{
			name:    "figure",
			command: "figure 0.25 0.25",
			op:      &painter.CrossFigure{CentralPoint: image.Point{X: 200, Y: 200}},
		},
		{
			name:    "move",
			command: "move 0.125 0.125",
			op:      &painter.MoveOperation{X: 100, Y: 100},
		},
		{
			name:    "update",
			command: "update",
			op:      painter.UpdateOp,
		},
		{
			name:    "invalid command",
			command: "invalidcommand",
			op:      nil,
		},
		{
			name:    "not enough args bigrect",
			command: "bgrect 0.125 0.125 0",
			op:      nil,
		},
		{
			name:    "too many args bigrect",
			command: "bgrect 0.125 0.125 0 0 0",
			op:      nil,
		},
		{
			name:    "wrong args bigrect",
			command: "bgrect w a c d",
			op:      nil,
		},
		{
			name:    "not enough args figure",
			command: "figure 0.125",
			op:      nil,
		},
		{
			name:    "too many args figure",
			command: "figure 0.125 0.125 0 0 0",
			op:      nil,
		},
		{
			name:    "wrong args figure",
			command: "figure ah",
			op:      nil,
		},
		{
			name:    "not enough args move",
			command: "move 0.125",
			op:      nil,
		},
		{
			name:    "too many args move",
			command: "move 0.125 0.125 0 0 0",
			op:      nil,
		},
		{
			name:    "wrong args move",
			command: "move gs j",
			op:      nil,
		},
	}
	// Looping through the test cases.
	for _, tc := range tests {
		// Running subtest for each test case.
		t.Run(tc.name, func(t *testing.T) {
			// Creating an instance of the parser.
			parser := &Parser{}
			// Parsing the command using the parser.
			ops, err := parser.Parse(strings.NewReader(tc.command))
			// If the expected operation is nil, we expect an error.
			if tc.op == nil {
				assert.Error(t, err)
			} else { // Otherwise, we expect no errors.
				require.NoError(t, err)
				// Asserting that the type of the parsed operation is the same as the expected operation.
				assert.IsType(t, tc.op, ops[1])
				// Asserting that the parsed operation is equal to the expected operation.
				assert.Equal(t, tc.op, ops[1])
			}
		})
	}
}

// Test_parse_func tests the Parse function of the parser.
func Test_parse_func(t *testing.T) {

	// A list of test cases consisting of commands to be parsed by the parser and their expected operations
	tests := []struct {
		name    string            // The name of the test case
		command string            // The command to be parsed
		op      painter.Operation // The expected operation to be returned by the parser
	}{
		{
			name:    "white fill",
			command: "white",
			op:      painter.OperationFunc(painter.WhiteFill), // The expected operation is a function call to WhiteFill
		},
		{
			name:    "green fill",
			command: "green",
			op:      painter.OperationFunc(painter.GreenFill), // The expected operation is a function call to GreenFill
		},
		{
			name:    "reset screen",
			command: "reset",
			op:      painter.OperationFunc(painter.Reset), // The expected operation is a function call to Reset
		},
	}

	// Create a new parser
	parser := &Parser{}

	// Iterate through the list of test cases
	for _, tc := range tests {
		// Run a sub-test for each test case
		t.Run(tc.name, func(t *testing.T) {
			// Parse the command using the parser
			ops, err := parser.Parse(strings.NewReader(tc.command))

			// Assert that there are no errors
			require.NoError(t, err)

			// Assert that the number of operations returned is 1
			require.Len(t, ops, 1)

			// Assert that the type of the first operation is the expected type
			assert.IsType(t, tc.op, ops[0])

		})
	}
}
