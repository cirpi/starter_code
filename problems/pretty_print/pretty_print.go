package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	MOD = int(1e9 + 7)
)

// Commonly used utility functions
func sum[T Number](ar []T) T {
	var ans T = 0
	for i := range ar {
		ans += ar[i]
	}
	return ans
}

func gcd[T Number](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm[T Number](a, b T) T {
	return (a * b) / gcd(a, b)
}

func abs[T Signed](x T) T {
	if x < 0 {
		return x * -1
	}
	return x
}
func mod[T Number](a, mod T) T {
	res := a % mod
	if res < 0 {
		res += mod
	}
	return res
}

func modAdd[T Number](a, b, Mod T) T {
	return mod(a+b, Mod)
}

func modMul[T Number](a, b, Mod T) T {
	return mod(a*b, Mod)
}

func modExp[T Number](base, exp, mod T) T {
	var result T = 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

func pow[T Precision](base, exp T) T {
	return T(math.Pow(float64(base), float64(exp)))
}

func upperBound[T Search](target T, ar []T) int {
	l, r := 0, len(ar)-1
	ans := -1
	for l <= r {
		m := l + (r-l)/2
		if ar[m] <= target {
			l = m + 1
		} else {
			ans = m
			r = m - 1
		}
	}
	return ans
}

func lowerBound[T Search](target T, ar []T) int {
	l, r := 0, len(ar)-1
	ans := -1
	for l <= r {
		m := l + (r-l)/2
		if ar[m] < target {
			l = m + 1
		} else {
			ans = m
			r = m - 1
		}
	}
	return ans
}

// Start
func main() {
	words := []string{
		"aaa",
		"bbb",
		"c",
		"d",
		"ee",
		"ff",
		"ggggggg",
	}
	line_length := 11
	// ans := solve(len(words)-1, line_length, words)
	ans := dpSolution(words, line_length)
	println(ans)
	// getLen(0, 1, words, line_length)
	flush()

}

func solve(ind, l int, words []string) int {
	if ind < 0 {
		return 0
	}
	ans := math.MaxInt
	for i := ind; i >= 0; i-- {
		if isPossible, length := getLen(i, ind, words, l); isPossible {
			ans = min(ans, solve(i-1, l, words)+length)
		}
	}
	return ans
}

func dpSolution(words []string, line_length int) int {
	dp := make([]int, len(words)+1)
	dp[0] = 0
	n := len(words)
	for ind := 1; ind <= n; ind++ {
		ans := math.MaxInt
		for i := ind; i > 0; i-- {
			if isPossible, length := getLen(i-1, ind-1, words, line_length); isPossible {
				ans = min(ans, dp[i-1]+length)
			}
		}
		dp[ind] = ans
	}
	return dp[n]
}

func getLen(i, j int, words []string, l int) (bool, int) {
	length := 0
	space := 0
	for ind := 0; ind <= (j - i); ind++ {
		length += len(words[i+ind])
		if ind > 0 {
			space++
		}
		if length+space > l {
			return false, 0
		}
	}
	// println(pow((l - (length + space)), 2))
	return true, pow((l - (length + space)), 2)
}

// Interfaces for my convenience
type Signed interface {
	int | int64 | int32
}
type Number interface {
	int | int64 | int32 | uint | uint64 | uint32
}

type Precision interface {
	Number | float64
}
type Search interface {
	Number
}

// IO
var RW bufio.ReadWriter
var IN = os.Stdin
var OUT = os.Stdout

func open(path string) *os.File {
	file, er := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0777)
	if er != nil {
		log.Fatal("open", er)
		return nil
	}
	return file
}

func initIO(in, out any) {
	if in != nil {
		val, ok := in.(string)
		if !ok {
			log.Fatal("initIO")
		}
		IN = open(val)
	}
	if out != nil {
		val, ok := out.(string)
		if !ok {
			log.Fatal("initIO")
		}
		OUT = open(val)
	}
}
func initReader(in, out any) {
	initIO(in, out)
	RW = *bufio.NewReadWriter(bufio.NewReader(IN), bufio.NewWriter(OUT))
}

func init() {
	initReader(nil, nil)
}

func flush() {
	RW.Flush()
}

// Reading, parsing, and assigning integer(s)
func readInt[T Number]() T {
	str, er := RW.ReadString('\n')
	if er != nil {
		log.Fatal("readInit", er)
	}
	str = strings.TrimSpace(str)
	val, er := strconv.Atoi(str)
	if er != nil {
		log.Fatal("readInit", er)
	}
	return T(val)
}

func readInts[T Number](addrs ...*T) []T {
	str, er := RW.ReadString('\n')
	if er != nil {
		log.Fatal("readInts", er)
	}
	str = strings.TrimSpace(str)
	ar := strings.Fields(str)
	vals := make([]T, len(ar))
	for i := range ar {
		val, er := strconv.Atoi(ar[i])
		if er != nil {
			log.Fatal("readInts", er)
		}
		vals[i] = T(val)
	}
	if len(addrs) > 0 {
		assign(vals, addrs...)
	}
	return vals
}

func assign[T any](ar []T, addrs ...*T) {
	if len(ar) < len(addrs) {
		log.Fatal("assign")
	}
	for i := range addrs {
		*addrs[i] = ar[i]
	}
}

// syntactic sugars for print functions
func println(args ...any) {
	fmt.Println(args...)
}

func print(args ...any) {
	fmt.Print(args...)
}

func printf(format string, args ...any) {
	fmt.Printf(format, args...)
}
