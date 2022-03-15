/*
	这是一个文档注释
*/
package commentary

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	这是顶级声明
*/

// 这是写评论的函数的注释
func Write() {
	fmt.Println(strings.Repeat(strconv.FormatInt(15, 16), 5))
}
