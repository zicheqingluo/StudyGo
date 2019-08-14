package main

import "fmt"


func main(){
	var (
		a = 5
		b = 2
	)
	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	b--
	fmt.Println(a ,b)

	//关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	//逻辑运算符
	//如果年龄大于18岁，并且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("苦逼")
	} else {
		fmt.Println("不用上班")
	}
	// 如果年龄小于18岁 或者 年龄大于60岁
	if age < 18 || age > 60{
		fmt.Println("不用上班")
	} else {
		fmt.Println("苦逼")
	}

	//not 取反，原来的真即为假，假即为真
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	//位运算：针对的是二进制数
	//5的二进制表示：101
	//2的二进制表示：10

	//&：按位与 （两位均为1才为1）
	fmt.Println(5 & 2) //000
	// |:按位或（两位有1个为1就为1）
	fmt.Println(5 | 2) // 111
	// ^:按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) // 111
	// <<:将二进制位左移指定位数
	fmt.Println(5 << 1)  // 1010 => 10
	fmt.Println(1 << 10) // 10000000000 => 1024
	// >>:将二进制位右移指定的位数
	fmt.Println(5 >> 2)  // 1 =>1
	var m = int8(1)      // 只能存8位
	fmt.Println(m << 10) // 10000000000

		// 192.168.1.1
	// 权限 文件操作会讲位运算实际的应用
	// 0644
	// 赋值运算符， 用来给变量赋值的

	var x int
	x = 10
	x += 1 // x = x + 1
	x -= 1 // x = x - 1
	x *= 2 // x = x * 2
	x /= 2 // x = x / 2
	x %= 2 // x = x % 2

	x <<= 2 // x = x << 2
	x &= 2  // x = x & 2
	x |= 3  // x = x | 3
	x ^= 4  // x = x ^ 4
	x >>= 2 // x = x >> 2

	fmt.Println(x)
}