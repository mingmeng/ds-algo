package string

import "testing"

func Test_replaceWords(t *testing.T) {
	// ["a", "aa", "aaa", "aaaa"]
	//"a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
	dic := []string{"a", "aa", "aaa", "aaaa"}
	sen := "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
	res := replaceWords(dic, sen)
	t.Log(res)
}
