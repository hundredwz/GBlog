package util

func IfMetaIn(mId int, mIds []int) bool {
	for _, i := range mIds {
		if mId == i {
			return true
		}
	}
	return false
}

func ReGenPostMap(postForm map[string][]string) map[string]interface{} {
	form := make(map[string]interface{})
	for key, value := range postForm {
		if len(value) == 0 {
			continue
		}
		if len(value) == 1 {
			form[key] = value[0]
		} else {
			form[key] = value
		}
	}
	return form
}
