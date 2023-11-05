package misc

func GetIndexesFromList[T any](arr []T) []int {
	var res []int
	for i := range arr {
		res = append(res, i)
	}

	return res
}
