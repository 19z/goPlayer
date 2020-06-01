package helper

func ArrayFilter(array []string) []string {
	var _array []string
	for _, _s := range array {
		if len(_s) > 0 {
			_array = append(_array, _s)
		}
	}
	return _array;
}

func ArraySpilt(array []string, start int, end int) []string {
	l := len(array)
	if end < 0 {
		end = l - end - 2;
	} else if (end == 0) {
		end = l
	}
	return array[start:end]
}

func ArrayEnd(array []string) string {
	return array[len(array)-1];
}

func InArray(array []string, val string) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}
