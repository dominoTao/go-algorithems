package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

type Node struct {
	ID  string
	PID string
}
// 现在数据库有一张表，用来存储一个多叉树，id为主键，pid 表示父节点的 id，已知 "-1" 表示根节点，现在要求打印出从根节点到每个子节点的路径（可以是无序的）。
func TestName(t *testing.T) {

	nodes := []Node{
		{
			"A",
			"-1",
		},
		{
			"A-1",
			"A",
		},
		{
			"A-2",
			"A",
		},
		{
			"A-3",
			"A",
		},
		{
			"A-2-1",
			"A-2",
		},
		{
			"A-2-2",
			"A-2",
		},
		{
			"A-2-3",
			"A-2",
		},
	}

	printPath(nodes)

	fmt.Println(nodes)
}

func printPath(n []Node){
	for _, v := range n {
		path := ""
		for {
			if v.PID == "-1" {
				path = fmt.Sprintf("%s%s%s", "/", v.ID, path)
				fmt.Println(path)
				break
			}else {
				path = fmt.Sprintf("%s%s%s", "/", v.ID, path)
				pre := prefix(v.PID, n)
				v = pre
			}
		}
	}
}

func prefix(id string, nodes []Node) Node {

	for _, v := range nodes {
		if v.ID == id {
			return v
		}
	}
	return Node{}
}
// 有一个字符串数组，每个字符串都只包含小写字母，现在需要找到两个长度相乘最大的字符串，并且两个字符串不能有相同的字母，如果没有满足这个条件的结果，返回0
func TestDisplaceMent(t *testing.T) {
	s := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(s)
	slice := genBinSlice(s)
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if slice[i] & slice[j] == 0 {
				t := len(s[i]) * len(s[j])
				if max < t {
					max = t
				}
			}
		}
	}
	fmt.Println(max)
}

func genBinSlice(s []string) []int {
	l := len(s)
	var temp []int
	for i := 0; i < l; i++ {
		str := s[i]
		var tt int
		for j := 0; j < len(str); j++ {
			i2 := 1 << (str[j] - 'a')
			tt |= i2
		}
		temp = append(temp, tt)
		fmt.Printf("%b\n", tt)
	}
	return temp
}

// 有一个字符串数组，每个字符串都只包含小写字母，现在需要找到两个长度相乘最大的字符串，并且两个字符串不能有相同的字母，如果没有满足这个条件的结果，返回0
func TestString(t *testing.T) {
	s := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	//s := []string{"a","aa","aaa","aaaa"}
	max := 0
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			count ++
			result := isContain(s[i], s[j])
			if result > max {
				max = result
			}
		}
	}
	fmt.Println(max)
	fmt.Println(count)
}

func isContain(a, b string) int {
	set1 := make(map[int32]struct{})
	set2 := make(map[int32]struct{})
	for _, v := range a {
		set1[v] = struct{}{}
	}
	set1Len := len(set1)
	for _, v := range b {
		set2[v] = struct{}{}
	}
	for v, _ := range set2 {
		set1[v] = struct{}{}
	}

	if set1Len + len(set2) == len(set1) {
		return len(a) * len(b)
	}
	return 0
}

// 从上到下找到最短路径（n个数字之和最小,n为矩阵的行数），可以从第一行中的任何元素开始，只能往下层走，同时只能走向相邻的节点，
//例如图中第一排 2 只能走向 第二排的 7、3；第二排的 7 可以走向第三排的 6、2、9
// // | 5    | 8    | 1    | 2    |
// // | 4    | 1    | 7    | 3    |
// // | 3    | 6    | 2    | 9    |
// //
// // Input: [
// //     [5, 8, 1, 2],
// //     [4, 1, 7, 3],
// //     [3, 6, 2, 9]
// // ]
// // Output: 4

func TestShortest(t *testing.T) {
	matrix := [][]int{
		{5, 8, 1, 2},
		{4, 1, 7, 3},
		{3, 6, 2, 9},
	}
	result := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		cols := matrix[i]
		if i == 0 {
			result = append(result, cols...)
			continue
		}
		for idx, _ := range cols {
			min := 0
			if idx == 0 {
				min = minOf2(cols[idx], cols[idx+1])
			}else if idx == len(cols)-1 {
				min = minOf2(cols[idx], cols[idx-1])
			} else {
				min = minOf3(cols[idx-1], cols[idx], cols[idx+1])
			}
			result[idx] += min
		}


	}
	theMin := minOfSlice(result)
	fmt.Println(theMin)
	fmt.Println(matrix)
}
func minOf2(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func minOf3(a, b, c int) int {
	min := 0
	if a > b {
		min = b
	} else {
		min = a
	}
	if min > c {
		min = c
	}
	return min
}

func minOfSlice(a []int) int {
	min := 1<<63 - 1
	fmt.Printf("max maxii   = %+v\n", min)
	for _, v := range a {
		if min > v {
			min = v
		}
	}
	return min
}


func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		time.Sleep(1*time.Second)
		fmt.Print(v)
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}
func main11() {
	fmt.Println("main", GoID())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, GoID())
		}()
	}
	wg.Wait()

}
func main() {
	a := []int{5, 3, 1, 5, 4}
	b := []int{5, 3}

	c := make(map[int]struct{})
	for _, v := range b {
		c[v] = struct{}{}
	}

	for i, v := range a {
		if _, ok := c[v]; ok {
			fmt.Println(i)
		}
	}

	fmt.Println(a)
	fmt.Println(b)
}
func main1() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	// 后一半数组
	go sum(s[:len(s)/2], c)
	// 前一半数组
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)  // -5 17 12  或者  17 -5 12
}

func Demo() {
	var num int
	var wg sync.WaitGroup  // 避免主线程提早退出
	var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			fmt.Println(i, "hello world ", num)
			mu.Lock()
			num++
			mu.Unlock()
			wg.Done()
		}(&wg, &mu)
	}
	wg.Wait()
}

func Demo2() {
	var num int
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i, "hello world ", num)
			num++
		}()
	}
	time.Sleep(2 * time.Second)
}