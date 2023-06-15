package compare

import "reflect"

func StringSliceByDeepEqual(sl1, sl2 []string) bool {
	return reflect.DeepEqual(sl1, sl2)
}

func StringSliceCustom1(sl1, sl2 []string) bool {
	equal := true
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			equal = false
			break
		}
	}
	return equal
}
