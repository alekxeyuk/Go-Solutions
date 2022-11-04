package binarysearch

func SearchInts(list []int, key int) int {
	size := len(list)
	cursor := size / 2

	for size > 0 {
		v := list[cursor]
		switch {
		case v == key:
			return cursor
		case cursor == 0 || cursor == size:
			return -1
		case v < key:
			oldC := cursor
			cursor += (size - cursor) / 2
			if oldC == cursor {
				return -1
			}
		case v > key:
			size = cursor
			cursor /= 2
		}
	}
	return -1
}
