package yacht

var cats = map[string]int{"ones": 1, "twos": 2, "threes": 3, "fours": 4, "fives": 5, "sixes": 6}

func Score(dice []int, category string) int {
	switch category {
	case "ones", "twos", "threes", "fours", "fives", "sixes":
		what := cats[category]
		return what * counter(&dice)[what]
	case "full house":
		fs, ss := false, false
		for _, v := range counter(&dice) {
			if v == 3 {
				fs = true
			}
			if v == 2 {
				ss = true
			}
		}
		if !(fs && ss) {
			return 0
		}
		fallthrough
	case "choice":
		return generic(&dice, func(i int) int { return i })
	case "four of a kind":
		for c, v := range counter(&dice) {
			if v >= 4 {
				return c * 4
			}
		}
		return 0
	case "yacht":
		if counter(&dice)[dice[0]] == 5 {
			return 50
		}
		return 0
	case "big straight", "little straight":
		ones := true
		six := false
		sum := 0
		for c, v := range counter(&dice) {
			if v != 1 {
				ones = false
			}
			if c == 6 {
				six = true
			}
			sum += c
		}
		if ones {
			if category == "big straight" && six && sum == 20 {
				return 30
			} else if category == "little straight" && !six && sum == 15 {
				return 30
			}
			return 0
		}
		return 0
	default:
		return 0
	}
}

func counter(in *[]int) map[int]int {
	counter := map[int]int{}
	for _, c := range *in {
		counter[c] += 1
	}
	return counter
}

func generic(in *[]int, fc func(int) int) (i int) {
	for _, c := range *in {
		i += fc(c)
	}
	return i
}
