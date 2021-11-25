package utils

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
)
func ZipStr(sb []byte) (content string) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(sb)
	w.Close()
	return b.String()
}

func UnzipStr(bs string) (s string) {
	var b bytes.Buffer
	b.WriteString(bs)
	r, err := zlib.NewReader(&b)
	if err != nil {
		fmt.Println(" err : ", err)
	}
	defer r.Close()
	bStr, err := ioutil.ReadAll(r)
	str := string(bStr[:])
	if err != nil {
		fmt.Println(" err : ", err)
	}
	return str
}