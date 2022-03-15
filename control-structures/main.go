package main

import (
	"fmt"
	"log"
	"os"
)

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func sumArr() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func openFile(fileName string) error {

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	_, err = f.Stat()
	if err != nil {
		f.Close()
		return err
	}
	return nil
}
func rangeProcessString() {
	for pos, char := range "日本\x80語" {
		fmt.Printf("字符 %#U 字节位置开始于 %d\n", char, pos)
	}
}

func main() {
	x := 100
	y := 101

	fmt.Printf("x 和 y 中较大的数是: %d \n", max(x, y))

	if err := openFile("./test.txt"); err != nil {
		log.Print(err)
	}
	fmt.Printf("0 - 10 的和是 %d \n", sumArr())
	rangeProcessString()
}
