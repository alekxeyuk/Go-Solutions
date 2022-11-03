package minesweeper

var neigboursShift = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	bLen := len(board)
	if bLen == 0 {
		return []string{}
	}
	rLen := len(board[0])
	newBoard := make([]string, bLen)
	newRow := make([]rune, rLen)
	for i, row := range board {
		copy(newRow, []rune(row))
		for x, c := range row {
			if c == '*' {
				continue
			} else {
				neigboursCount := 0
				for _, shift := range neigboursShift {
					sY := i + shift[0]
					sX := x + shift[1]
					if sY >= 0 && sY < bLen && sX >= 0 && sX < rLen {
						if board[sY][sX] == '*' {
							neigboursCount++
						}
					}
				}
				if neigboursCount != 0 {
					newRow[x] = rune(neigboursCount + 48)
				}
			}
		}
		newBoard[i] = string(newRow)
	}
	return newBoard
}
