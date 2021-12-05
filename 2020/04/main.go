package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

type passport struct {
	byr    string
	iyr    string
	eyr    string
	hgt    string
	hcl    string
	ecl    string
	pid    string
	cid    string
	reason string
}

func main() {
	passports := parse(utils.Load(2020, "04-real").String())
	numValid := 0
	for _, p := range passports {
		if p.isValid() {
			numValid++
		}
		fmt.Println(p.String())
	}
	fmt.Printf("%d valid\n", numValid)
}

func parse(s string) []passport {
	retVal := make([]passport, 0)

	splitted := strings.Split(s, "\n\n")
	for _, part := range splitted {
		thisPass := passport{}
		cleanPart := strings.ReplaceAll(part, "\n", " ")
		splitterPass := strings.Split(cleanPart, " ")
		for _, passPart := range splitterPass {
			parseKeyValue(passPart, &thisPass)
		}
		retVal = append(retVal, thisPass)
	}

	return retVal
}

var reHCL = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var rePIC = regexp.MustCompile(`^[0-9]{9}$`)
var reSize = regexp.MustCompile(`^[0-9]+`)
var validECL = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func (p *passport) byrValid() bool {
	if p.byr == "" {
		p.reason = "empty byr"
		return false
	}
	n, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}
	if n < 1920 || n > 2002 {
		p.reason = fmt.Sprintf("out of range: %d (1920 < 2002)", n)
		return false
	}
	return true
}

func (p *passport) iyrValid() bool {
	if p.iyr == "" {
		p.reason = "empty iyr"
		return false
	}
	n, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}
	if n < 2010 || n > 2020 {
		p.reason = fmt.Sprintf("out of range: %d (2010 < 2020)", n)
		return false
	}

	return true
}

func (p *passport) eyrValid() bool {
	if p.eyr == "" {
		p.reason = "empty eyr"
		return false
	}
	n, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}
	if n < 2020 || n > 2030 {
		p.reason = fmt.Sprintf("out of range: %d (2020 < 2030)", n)
		return false
	}
	return true
}

func (p *passport) hgtValid() bool {
	if p.hgt == "" {
		p.reason = "empty hgt"
		return false
	}
	if strings.HasSuffix(p.hgt, "in") {
		size := reSize.FindAllString(p.hgt, -1)
		if len(size) != 1 {
			return false
		}
		sizeInt, err := strconv.Atoi(size[0])
		if err != nil {
			return false
		}
		if sizeInt < 59 || sizeInt > 76 {
			p.reason = fmt.Sprintf("out of range: %s (59 < 76)", p.hgt)
			return false
		}
	} else if strings.HasSuffix(p.hgt, "cm") {
		size := reSize.FindAllString(p.hgt, -1)
		if len(size) != 1 {
			return false
		}
		sizeInt, err := strconv.Atoi(size[0])
		if err != nil {
			return false
		}
		if sizeInt < 150 || sizeInt > 193 {
			p.reason = fmt.Sprintf("out of range: %s (150 < 193)", p.hgt)
			return false
		}
	} else {
		p.reason = fmt.Sprintf("no valid unit: %s", p.hgt)
		return false
	}

	return true

}

func (p *passport) hclValid() bool {
	if p.hcl == "" {
		p.reason = "empty hcl"
		return false
	}
	if !reHCL.MatchString(p.hcl) {
		p.reason = fmt.Sprintf("invalid hcl %s", p.hcl)
		return false
	}
	return true
}

func (p *passport) eclValid() bool {
	if p.ecl == "" {
		p.reason = "empty ecl"
		return false
	}
	foundECL := false
eclLoop:
	for _, v := range validECL {
		if v == p.ecl {
			foundECL = true
			break eclLoop
		}
	}
	if !foundECL {
		p.reason = fmt.Sprintf("ecl not found: %s", p.ecl)
		return false
	}
	return true
}

func (p *passport) pidValid() bool {
	if p.pid == "" {
		p.reason = "empty pid"
		return false
	}
	if !rePIC.MatchString(p.pid) {
		p.reason = fmt.Sprintf("invalid pid: %s", p.pid)
		return false
	}
	return true
}

func (p *passport) isValid() bool {
	if !p.byrValid() {
		return false
	}
	if !p.iyrValid() {
		return false
	}
	if !p.eyrValid() {
		return false
	}
	if !p.hgtValid() {
		return false
	}
	if !p.hclValid() {
		return false
	}
	if !p.eclValid() {
		return false
	}
	if !p.pidValid() {
		return false
	}

	return true
}

func (p passport) String() string {
	if p.reason != "" {
		return fmt.Sprintf("byr:%4s iyr:%4s eyr:%4s hgt:%6s hcl:%7s ecl:%7s pid:%10s -> %s", p.byr, p.iyr, p.eyr, p.hgt, p.hcl, p.ecl, p.pid, p.reason)
	}
	return fmt.Sprintf("byr:%4s iyr:%4s eyr:%4s hgt:%6s hcl:%7s ecl:%7s pid:%10s", p.byr, p.iyr, p.eyr, p.hgt, p.hcl, p.ecl, p.pid)
}

func parseKeyValue(s string, pass *passport) {
	splitted := strings.Split(s, ":")
	if len(splitted) == 2 {
		key := splitted[0]
		value := splitted[1]
		switch key {
		case "byr":
			pass.byr = value
		case "iyr":
			pass.iyr = value
		case "eyr":
			pass.eyr = value
		case "hgt":
			pass.hgt = value
		case "hcl":
			pass.hcl = value
		case "ecl":
			pass.ecl = value
		case "pid":
			pass.pid = value
		case "cid":
			pass.cid = value
		}
	}
}
