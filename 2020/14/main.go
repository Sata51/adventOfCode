package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	data := utils.Load("real").ToStringSlice()

	fmt.Println(part1(data))
	fmt.Println(part2(data))

}

func to36bin(s string) string {
	var i int64
	i, _ = strconv.ParseInt(s, 10, 36)

	binStr := strconv.FormatInt(i, 2)
	var buffer bytes.Buffer
	padding := 36 - len(binStr)
	for padding != 0 {
		buffer.WriteRune('0')
		padding--
	}
	buffer.WriteString(binStr)
	return buffer.String()
}

func applyMask(bin, mask string) string {
	var buffer bytes.Buffer

	for i, n := range mask {
		if n != 'X' {
			buffer.WriteRune(n)
			continue
		}
		buffer.WriteByte(bin[i])
	}
	return buffer.String()
}

func part1(lines []string) (answer int) {
	mem := map[string]string{}
	mask := ""
	for _, line := range lines {
		if line == "" {
			continue
		}
		data := strings.Split(line, "=")
		if strings.HasPrefix(data[0], "ma") {
			mask = data[1][1:]
			continue
		}
		address := data[0][4 : len(data[0])-2]
		value := data[1][1:]

		masked := applyMask(to36bin(value), mask)
		mem[address] = masked
	}
	for _, v := range mem {
		val, err := strconv.ParseInt(v, 2, 64)
		if err != nil {
			panic(err)
		}
		answer += int(val)
	}

	return
}

func applyAddressMask(address, mask string) string {
	var buffer bytes.Buffer
	for i, n := range mask {
		if n == 'X' || n == '1' {
			buffer.WriteRune(n)
			continue
		}
		buffer.WriteByte(address[i])
	}
	return buffer.String()
}

func generateAddresses(addr string) []string {
	addresses := []*bytes.Buffer{{}}
	direction := true
	for _, c := range addr {
		if c == 'X' {
			size := len(addresses) - 1
			for ; 0 <= size; size-- {
				var cp bytes.Buffer
				cp.WriteString(addresses[size].String())
				addresses = append(addresses, &cp)
			}
			for i, address := range addresses {
				if direction {
					if i%2 == 0 {
						address.WriteRune('0')
					} else {
						address.WriteRune('1')
					}
				} else {
					if i%2 == 0 {
						address.WriteRune('1')
					} else {
						address.WriteRune('0')
					}
				}
			}
			direction = !direction
		} else {
			for _, address := range addresses {
				address.WriteRune(c)
			}
		}
	}
	var res []string
	for _, a := range addresses {
		res = append(res, a.String())
	}
	return res
}

func part2(lines []string) (answer int) {
	mem := map[string]string{}
	mask := ""
	for _, line := range lines {
		if line == "" {
			continue
		}
		data := strings.Split(line, "=")
		if strings.HasPrefix(data[0], "ma") {
			mask = data[1][1:]
			continue
		}
		address := data[0][4 : len(data[0])-2]
		value := data[1][1:]

		masked := applyAddressMask(to36bin(address), mask)
		addresses := generateAddresses(masked)
		for _, m := range addresses {
			mem[m] = value
		}
	}
	for _, v := range mem {
		answer += utils.MustParseInt(v)
	}
	return
}
