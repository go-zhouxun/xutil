package xstring

import "bytes"

func StringJoin(strList ...string) string {
	b := bytes.Buffer{}
	for _, str := range strList {
		b.WriteString(str)
	}
	return b.String()
}
