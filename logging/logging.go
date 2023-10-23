package logging

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NoahOnFyre/gengine/color"
)

var mainColor string = color.GreenBg

func SetMainColor(str string) {
	mainColor = str
}

func Log(msg ...any) {
	msg = append(Ar(Prefix(0)), msg...)
	fmt.Println(msg...)
}

func Warn(msg ...any) {
	msg = append(Ar(Prefix(1)), msg...)
	fmt.Println(msg...)
}

func Error(msg ...any) {
	msg = append(Ar(Prefix(2)), msg...)
	fmt.Println(msg...)
}

func Print(msg ...any) {
	fmt.Println(msg...)
}

func PrintR(msg ...any) {
	fmt.Print(msg...)
}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	PrintR(msg)
	scanner.Scan()
	return scanner.Text()
}

func Ar(msg string) []any {
	return []any{msg}
}

func Clear() {
	PrintR("\033[H\033[2J")
}
