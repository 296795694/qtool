package encrypt

import (
	baseMd5 "crypto/md5"
	"fmt"
	"io"
)

// Md5 md5加密
func Md5(str string) (md5str string) {
	w := baseMd5.New()
	io.WriteString(w, str)
	md5str = fmt.Sprintf("%x", w.Sum(nil))
	return
}
