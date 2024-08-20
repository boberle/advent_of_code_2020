package day13

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Timestamp int

type BusInfo struct {
	id     int
	offset int
}

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	t, busInfo := parseFile(fh)
	fmt.Printf("Earliest bus * waiting time (part 1): %d\n", findEarliestBus(busInfo, t))
	fmt.Printf("Earliest timestamp (part 2): %d\n", solve(busInfo))

}

func parseFile(reader io.Reader) (Timestamp, []BusInfo) {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	timestamp, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	data := strings.Split(scanner.Text(), ",")
	busInfo := []BusInfo{}
	for i, d := range data {
		if d != "x" {
			id, err := strconv.Atoi(d)
			if err != nil {
				panic(err)
			}
			busInfo = append(busInfo, BusInfo{id: id, offset: i})
		}
	}

	return Timestamp(timestamp), busInfo
}

func findEarliestBus(busInfo []BusInfo, t Timestamp) int {
	minWaitingTime := 0
	minId := 0
	for i, bi := range busInfo {
		waitingTime := (int(t)/bi.id+1)*bi.id - int(t)
		if i == 0 || waitingTime < minWaitingTime {
			minWaitingTime = waitingTime
			minId = bi.id
		}
	}
	return minId * minWaitingTime
}

func inv_mod(a, m int) int {
	for i := 0; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	panic("inv mod not found")
}

func solve(busInfo []BusInfo) int {
	M := 1
	for _, bi := range busInfo {
		M *= bi.id

	}

	x0 := 0
	for _, bi := range busInfo {
		a := modulo(bi.id-bi.offset, bi.id)
		Mi := M / bi.id
		x0 += a * Mi * inv_mod(Mi, bi.id)
	}

	x := x0 % M
	return x
}

func modulo(a, b int) int {
	return (a%b + b) % b
}
