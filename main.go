package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
	"math"
	"strconv"
)

func main() {
	date1 := ReadEntry()
	date2 := ReadEntry()
	fmt.Println(TimeBetweenTwoDate(date1, date2))
}

func ReadEntry() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Date: ")
	ti, _ := reader.ReadString('\n')
	text := strings.Replace(ti, "\n", "", -1)
	return text
}

func TimeBetweenTwoDate(time1 string, time2 string) string {
	timeFormat := "2006-01-02"
	t1, _ := time.Parse(timeFormat, time1)
	t2, _ := time.Parse(timeFormat, time2)

	duration := t2.Sub(t1)

	arrondi := math.Floor((duration.Hours() / 24) /7 )
	i := int(arrondi)

	weeksBetweenTwoDates := strconv.Itoa(i)

	return  weeksBetweenTwoDates + " week(s)"
}