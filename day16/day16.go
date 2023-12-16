package day16

import (
	"slices"
)

// N1 computes the results for Ex1 on the given input-file
func N1(in []string) int {
	board := parse(in)
	var touched = make([][][]int, len(board))
	for i := range touched {
		touched[i] = make([][]int, len(board[i]))
	}
	simulateLight(&board, &touched, 0, 0, right)

	fieldCount := 0
	for _, row := range touched {
		for _, f := range row {
			if len(f) != 0 {
				fieldCount++
			}
		}
	}

	return fieldCount
}

// N2 computes the results for Ex2 on the given input-file
func N2(in []string) int {
	board := parse(in)
	maxV := 0
	for i := range board {
		var touched = make([][][]int, len(board))
		for i := range touched {
			touched[i] = make([][]int, len(board[i]))
		}
		simulateLight(&board, &touched, i, 0, right)

		fieldCount := 0
		for _, row := range touched {
			for _, f := range row {
				if len(f) != 0 {
					fieldCount++
				}
			}
		}
		if fieldCount > maxV {
			maxV = fieldCount
		}

		touched = make([][][]int, len(board))
		for i := range touched {
			touched[i] = make([][]int, len(board[i]))
		}
		simulateLight(&board, &touched, i, len(board[i])-1, left)

		fieldCount = 0
		for _, row := range touched {
			for _, f := range row {
				if len(f) != 0 {
					fieldCount++
				}
			}
		}
		if fieldCount > maxV {
			maxV = fieldCount
		}
	}
	for i := range board[0] {
		var touched = make([][][]int, len(board))
		for i := range touched {
			touched[i] = make([][]int, len(board[i]))
		}
		simulateLight(&board, &touched, 0, i, down)

		fieldCount := 0
		for _, row := range touched {
			for _, f := range row {
				if len(f) != 0 {
					fieldCount++
				}
			}
		}
		if fieldCount > maxV {
			maxV = fieldCount
		}

		touched = make([][][]int, len(board))
		for i := range touched {
			touched[i] = make([][]int, len(board[i]))
		}
		simulateLight(&board, &touched, len(board)-1, i, up)

		fieldCount = 0
		for _, row := range touched {
			for _, f := range row {
				if len(f) != 0 {
					fieldCount++
				}
			}
		}
		if fieldCount > maxV {
			maxV = fieldCount
		}
	}

	return maxV
}

const (
	up    = 1
	down  = 2
	left  = 3
	right = 4
)

func simulateLight(board *[][]rune, touched *[][][]int, x int, y int, dir int) {
	if x < 0 || x >= len(*board) || y < 0 || y >= len((*board)[0]) {
		return
	}
	switch (*board)[x][y] {
	case '.':
		newX, newY := modCoordinates(x, y, dir)
		if slices.Contains((*touched)[x][y], dir) {
			return
		}
		(*touched)[x][y] = append((*touched)[x][y], dir)
		simulateLight(board, touched, newX, newY, dir)
	case '/':
		switch dir {
		case up:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x, y+1, right)
		case down:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x, y-1, left)
		case left:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x+1, y, down)
		case right:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x-1, y, up)
		}
	case '\\':
		switch dir {
		case up:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x, y-1, left)
		case down:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x, y+1, right)
		case left:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x-1, y, up)
		case right:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x+1, y, down)
		}
	case '|':
		switch dir {
		case up, down:
			newX, newY := modCoordinates(x, y, dir)
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, newX, newY, dir)
		case left, right:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x+1, y, down)
			simulateLight(board, touched, x-1, y, up)
		}
	case '-':
		switch dir {
		case up, down:
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, x, y+1, right)
			simulateLight(board, touched, x, y-1, left)
		case left, right:
			newX, newY := modCoordinates(x, y, dir)
			(*touched)[x][y] = append((*touched)[x][y], dir)
			simulateLight(board, touched, newX, newY, dir)
		}
	}
}

func modCoordinates(x int, y int, dir int) (newX int, newY int) {
	newX = x
	newY = y
	switch dir {
	case up:
		newX--
	case down:
		newX++
	case left:
		newY--
	case right:
		newY++
	}
	return
}

func parse(lines []string) (board [][]rune) {
	for _, l := range lines {
		var boardLine []rune
		for _, c := range l {
			boardLine = append(boardLine, c)
		}
		board = append(board, boardLine)
	}
	return
}
