package utils

func AssertEqual(origin interface{}, project interface{}) {
	if origin == project {
		return
	} else {
		panic("Assert Equal Faile")
	}
}
