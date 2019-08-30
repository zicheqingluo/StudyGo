/*
 * @Author: yangxiaokang
 * @Date: 2019-08-29 17:22:16
 * @Last Modified by: yangxiaokang
 * @Last Modified time: 2019-08-30 09:51:31
 */

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)                            //如同shell的$*
	fmt.Println(os.Args[0], os.Args[1], os.Args[2]) //如同shell $1 $2
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("序号%v 值%v \n", index, arg)  //os.args本质是字符串切片
		}
	}
}
