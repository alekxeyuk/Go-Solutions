package stateoftictactoe

import (
	"errors"
	"strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func checkRows(board []string, player string) bool {
	for _, row := range board {
		if strings.Count(row, player) == 3 {
			return true
		}
	}
	return false
}

func checkDiagonals(board []string, player rune) bool {
	winl, winr := 0, 0
	for i := 0; i < 3; i++ {
		if rune(board[i][i]) == player {
			winl++
		}
		if rune(board[i][2-i]) == player {
			winr++
		}
	}
	return winl == 3 || winr == 3
}

func transpose(board []string) ([]string, int, int) {
	xs, os := 0, 0
	transposed := make([]string, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transposed[i] += string(board[j][i])
			if board[j][i] == 'X' {
				xs++
			} else if board[j][i] == 'O' {
				os++
			}
		}
	}
	return transposed, xs, os
}

func StateOfTicTacToe(board []string) (State, error) {
	transposed, xs, os := transpose(board)
	xw := checkRows(board, "X") || checkRows(transposed, "X") || checkDiagonals(board, 'X')
	ow := checkRows(board, "O") || checkRows(transposed, "O") || checkDiagonals(board, 'O')

	if os > xs || xs > os+1 || xw && ow {
		return "", errors.New("invalid board")
	}

	if xw || ow {
		return Win, nil
	} else if xs+os == 9 {
		return Draw, nil
	} else {
		return Ongoing, nil
	}
}
