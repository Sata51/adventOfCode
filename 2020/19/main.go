package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

type Rule struct {
	contents string
	literal  bool
	needs    []int
}

func main() {
	input := utils.Load("real")

	fmt.Println("-----Part 1-----")
	fmt.Println("Result: ", part1(input.String()))
	fmt.Println("-----Part 2-----")
	fmt.Println("Result: ", part2(input.ToStringSlice()))
}

func part1(input string) int {
	s := strings.Split(input, "\n\n")
	rules, inputs := s[0], s[1]
	regex := regexp.MustCompile(parseRegex(strings.Split(rules, "\n")))

	cnt := 0
	for _, line := range strings.Split(inputs, "\n") {
		if regex.MatchString(line) {
			cnt++
		}
	}
	return cnt
}

func part2(input []string) int {
	type rule struct {
		val    string
		groups [][]string
	}

	rules := map[string]rule{}

	i := 0

	for ; ; i++ {
		line := input[i]

		if line == "" {
			i++
			break
		}

		tokens := strings.Split(line, ": ")
		k := tokens[0]

		if k == "8" {
			tokens[1] = "42 | 42 8"
		}

		if k == "11" {
			tokens[1] = "42 31 | 42 11 31"
		}

		if strings.HasPrefix(tokens[1], "\"") {
			rules[k] = rule{
				val: strings.Trim(tokens[1], "\""),
			}

			continue
		}

		r := rule{
			groups: [][]string{},
		}

		for _, str := range strings.Split(tokens[1], " | ") {
			r.groups = append(r.groups, strings.Split(str, " "))
		}

		rules[k] = r
	}

	seen := map[string]bool{}

	var iter func(key string, then []string, search string, index int)
	iter = func(key string, then []string, search string, index int) {
		rule := rules[key]

		if rule.val != "" {
			if index >= len(search) {
				return
			}

			if rule.val == search[index:index+1] {
				if len(then) > 0 {
					iter(then[0], then[1:], search, index+1)
				} else {
					if index == len(search)-1 {
						seen[search] = true
					}
				}
			}

			return
		}

		g1 := rule.groups[0]
		iter(g1[0], append(g1[1:], then...), search, index)

		if len(rule.groups) > 1 {
			g2 := rule.groups[1]
			iter(g2[0], append(g2[1:], then...), search, index)
		}
	}

	count := 0

	for ; i < len(input); i++ {
		iter(rules["0"].groups[0][0], rules["0"].groups[0][1:], input[i], 0)
		_, ok := seen[input[i]]
		if ok {
			count++
		}
	}

	return count
}

func parseRegex(lines []string) string {
	resolved, unresolved := make(map[int]Rule), make(map[int]Rule)
	for _, line := range lines {
		split := strings.Split(line, ": ")
		index, rule := utils.MustParseInt(split[0]), Rule{
			contents: split[1],
		}

		if rule.contents[0] == '"' {
			rule.contents = strings.Replace(rule.contents, "\"", "", 2)
			rule.literal = true
			resolved[index] = rule
		} else {
			rule.needs = parseUniqueNumbersFromString(strings.Replace(rule.contents, " | ", " ", -1))
			unresolved[index] = rule
		}
	}

	for len(unresolved) > 0 {
		for i, rule := range unresolved {
			if !containsAllKeys(resolved, rule.needs) {
				continue
			}

			// We can resolve this one!
			scanner := bufio.NewScanner(strings.NewReader(rule.contents))
			scanner.Split(bufio.ScanWords)
			var sb strings.Builder

			for scanner.Scan() {
				token := scanner.Text()
				if token == "|" {
					sb.WriteString("|")
					continue
				}

				otherRule := resolved[utils.MustParseInt(token)]
				sb.WriteString(otherRule.contents)
			}

			rule.contents = "(" + sb.String() + ")"

			resolved[i] = rule
			delete(unresolved, i)
		}
	}

	// All rules have been parsed
	return "^" + resolved[0].contents + "$"
}

func parseUniqueNumbersFromString(s string) []int {
	keys := make(map[int]bool)
	var list []int
	for _, n := range strings.Split(s, " ") {
		num := utils.MustParseInt(n)
		if _, value := keys[num]; !value {
			keys[num] = true
			list = append(list, num)
		}
	}
	return list
}

func containsAllKeys(haystack map[int]Rule, needle []int) bool {
	for _, search := range needle {
		if _, ok := haystack[search]; !ok {
			return false
		}
	}

	return true
}
