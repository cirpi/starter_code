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
	t := readInt[int]()
	for t > 0 {
		n := readInt[int]()
		str := readString()
		moves := 0
		for i := 0; i < n-1; i++ {
			if str[i] != str[i+1] {
				moves++
			}
		}
		if str[0] == '1' {
			moves++
		}
		println(moves)
		t--
	}
	flush()

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
func readString() string {
	str, er := RW.ReadString('\n')
	if er != nil {
		log.Fatal(er)
	}
	return strings.TrimSpace(str)
}

func readStrings() []string {
	str := readString()
	return strings.Fields(str)
}
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
