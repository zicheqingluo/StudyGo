/*
 * @Author: yangxiaokang
 * @Date: 2019-08-29 17:22:16
 * @Last Modified by: yangxiaokang
 * @Last Modified time: 2019-08-29 18:40:41
 */

package main
import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) >  0 {
		for index,arg := range os.Args {
			fmt.Printf("args[%d]=%v\n",index,arg)
		}
	}

}