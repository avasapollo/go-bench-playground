package compare

import "reflect"

func StringSliceByDeepEqual(sl1, sl2 []string) bool {
	return reflect.DeepEqual(sl1, sl2)
}

func StringSliceCustom1(sl1, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}

	return true
}

func StringSlice(sl1, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}

	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}

	return true
}
