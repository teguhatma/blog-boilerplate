package controller

import "strconv"

func convertToInt32(l, o string) (int32, int32, error) {
	limit, err := strconv.Atoi(l)
	if err != nil {
		return -1, -1, err
	}
	offset, err := strconv.Atoi(o)
	if err != nil {
		return -1, -1, err
	}

	a := int32(limit)
	b := int32(offset)

	return a, b, nil
}
