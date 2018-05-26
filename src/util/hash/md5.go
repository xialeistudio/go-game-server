// 哈希算法
package hash

import (
	"crypto/md5"
	"fmt"
	"io"
)

// md5算法
func Md5(reader io.Reader) string {
	m := md5.New()
	io.Copy(m, reader)
	return fmt.Sprintf("%x", m.Sum(nil))
}
