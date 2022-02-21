package wordy

import (
	"strconv"
	"strings"
)

type operand int

const (
	NONE operand = iota
	PLUS
	MINUS
	MUL
	DIV
)

type state struct {
	FIRST int64
	OP    operand
	SEC   int64
}

var operandTable = map[string]operand{
	"plus":       PLUS,
	"minus":      MINUS,
	"multiplied": MUL,
	"divided":    DIV,
}

func (s *state) compute() int64 {
	switch s.OP {
	case PLUS:
		return s.FIRST + s.SEC
	case MINUS:
		return s.FIRST - s.SEC
	case MUL:
		return s.FIRST * s.SEC
	case DIV:
		return s.FIRST / s.SEC
	default:
		return 0
	}
}

func Answer(question string) (int, bool) {
	tokens := strings.Split(question, " ")
	tokensLen := len(tokens)
	if tokensLen < 3 {
		return 0, false
	}

	tokens[tokensLen-1] = tokens[tokensLen-1][:len(tokens[tokensLen-1])-1]
	st := state{OP: NONE}
	f, err := strconv.ParseInt(tokens[2], 10, 64)
	if err != nil {
		return 0, false
	}
	st.FIRST = f
	for i := 3; i < tokensLen; i++ {
		if st.OP == NONE {
			op, e := operandTable[tokens[i]]
			if !e {
				return 0, false
			}
			st.OP = op
			if st.OP == MUL || st.OP == DIV {
				i++
			}
		} else {
			f, err := strconv.ParseInt(tokens[i], 10, 64)
			if err != nil {
				return 0, false
			}
			st.SEC = f
			st.FIRST = st.compute()
			st.OP = NONE
			st.SEC = 0
		}
	}

	if st.OP != NONE {
		return 0, false
	}
	return int(st.FIRST), true
}
