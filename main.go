package main

import (
	"fmt"
	"hook/parser"
	"os"
	"strconv"
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func main() {
	var grid [][]string
	var stacks []*Stack

	loc := Location{0, 0, &grid}
	vel := Velocity{1, 0}
	running := true

	var handle *os.File

	if len(os.Args) >= 2 {
		handle, _ = os.Open(os.Args[1])
	} else {
		panic("Please give a filename containing your fish program")
	}

	defer handle.Close()
	grid = parser.Tokenize(handle)

	Use(loc, vel, running, grid, stacks)

	stacks = append(stacks, NewStack())

	for running {
		sym := grid[loc.X][loc.Y]
		stack := stacks[len(stacks)-1]

		switch sym {
		case ";":
			running = false

		case "v":
			vel = Velocity{0, 1}
		case "^":
			vel = Velocity{0, -1}
		case "<":
			vel = Velocity{-1, 0}
		case ">":
			vel = Velocity{1, 0}

		case "[":
			new_stack := stack.Split()

			stacks = append(stacks, new_stack)
		case "]":
			previous_stack := stacks[len(stacks)-2]
			previous_stack.Join(stack)

			stacks = stacks[0 : len(stacks)-1]

		case "\\":
			if vel.X != 0 {
				vel.Y = vel.X * -1
				vel.X = 0
			} else if vel.Y != 0 {
				vel.X = vel.Y
				vel.Y = 0
			}

		case "r":
			stack.Reverse()
		case "n":
			fmt.Print(stack.Pop())
		case "d":
			fmt.Println(stack.Nodes)
		case "x":
			fmt.Println(len(stack.Nodes))

		default:
			i, err := strconv.Atoi(sym)
			if err == nil {
				stack.Push(i)
			}
		}

		loc.Move(vel)
	}
}
