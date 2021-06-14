package main

import (
	"bufio"
	"os"
	"strings"
)

func get_dimensions(handle *os.File) (x int, y int) {
	handle.Seek(0, 0)
	x, y = 0, 0

	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		y += 1
		line_length := len(scanner.Text())

		if line_length > x {
			x = line_length
		}
	}

	return
}

func read_line(handle *os.File, line int) string {
	handle.Seek(0, 0)
	scanner := bufio.NewScanner(handle)
	line_no := 0

	for scanner.Scan() {
		if line_no == line {
			return scanner.Text()
		}

		line_no++
	}

	return ""
}

func Tokenize(handle *os.File) [][]string {
	handle.Seek(0, 0)
	x_size, y_size := get_dimensions(handle)

	var ret [][]string

	for i := 0; i < y_size; i++ {
		row := strings.Split(read_line(handle, i), "")

		for len(row) < x_size {
			row = append(row, "")
		}

		ret = append(ret, row)
	}

	return ret
}
