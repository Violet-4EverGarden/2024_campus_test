package main

import "fmt"

// 小红的排列构造: (a[i] + b[i]) % n == c[i] % n
// n为偶数时必然不存在这样的三个排列：假设n=2k，根据求余运算的规律（如果n为偶数，那么m对n求余的结果r和m的奇偶性必定一致）：
// 假设排列c的前k个为奇数，后k个为偶数，那么a和b找不到满足要求的对应排列（奇+奇=偶、奇+偶=奇）：例如，c的前k个为奇数，那么a的前k个与b的前k个的
// 对应位数之和也必须为奇数（只能是奇+偶），但二者的后k个就不满足对应数字之和为偶数的特性了，因此它们的求余也必定不可能相同。
// 故，n为偶数时必然为false
// n为奇数时，只需满足c[i] = b[i] + a[i]，a从1取到n，b从n开始每次+1（真正的b需要对n取余），c也类似（需要对n求余）
func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Print(false)
		return
	}
	var a = make([]int, n+1)
	var b = make([]int, n+1)
	var c = make([]int, n+1)
	a[1] = 1
	b[1] = n
	c[1] = (b[1] + 1) % n
	for i := 2; i <= n; i++ {
		a[i] = i
		b[i] = (b[i-1] + 1) % n
		if b[i] == 0 {
			b[i] = n
		}
		c[i] = (c[i-1] + 2) % n
		if c[i] == 0 {
			c[i] = n
		}
		// 可以再判断一下
		if (a[i]+b[i])%n == c[i]%n {
			continue
		} else {
			fmt.Print(false)
			return
		}
	}
	fmt.Print("YES\n")
	fmt.Println(a[1:])
	fmt.Println(b[1:])
	fmt.Println(c[1:])
}

/*
3

YES
1 2 3
3 1 2
1 3 2
*/
