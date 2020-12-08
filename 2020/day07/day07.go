package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type bag struct {
	colour   string
	contents []content
}

type content struct {
	number     int
	contentBag *bag
}

func solvePartOne(input []string) int {
	bags := makeBagTree(input)
	target := bags["shiny gold bag"]
	containingBags := 0

	for _, bag := range bags {
		if find(bag, target) {
			containingBags++
		}
	}

	return containingBags
}

func solvePartTwo(input []string) int {
	bags := makeBagTree(input)
	target := bags["shiny gold bag"]

	total := findChildBags(target)
	return total
}

func makeBagTree(input []string) map[string]*bag {
	bags := make(map[string]*bag)

	// make map first
	for _, line := range input {
		parts := strings.Split(line, " contain ")
		if len(parts) < 2 {
			fmt.Errorf("cannot parse string '%s'", line)
		}

		// trim off trailing s to make life easier
		key := parts[0]
		if strings.HasSuffix(key, "s") {
			key = key[:len(key)-1]
		}

		bags[key] = &bag{
			colour:   key,
			contents: []content{},
		}
	}

	// now go through and add contents
	for _, line := range input {
		parts := strings.Split(line, " contain ")
		key := parts[0]
		if strings.HasSuffix(key, "s") {
			key = key[:len(key)-1]
		}

		b, ok := bags[key]
		if !ok {
			fmt.Errorf("bag was not added: '%s'", parts[0])
		}

		if parts[1] == "no other bags." {
			continue
		}

		contents := strings.Split(parts[1], ", ")
		for _, con := range contents {
			spaceIndex := strings.Index(con, " ")
			if spaceIndex == -1 {
				fmt.Errorf("unexpected bag string '%s'", con)
			}
			numStr := con[:strings.Index(con, " ")]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Errorf("Could not convert '%s' to int", numStr)
			}

			key = con[spaceIndex+1:]

			// the last bag will have a . after the key
			if strings.HasSuffix(key, ".") {
				key = key[:len(key)-1]
			}

			// a bag contains n bagS, so we need to trim the trailing S if it exists
			// in order to match our key
			if strings.HasSuffix(key, "s") {
				key = key[:len(key)-1]
			}

			c, ok := bags[key]
			if !ok {
				fmt.Errorf("could not find contained bag '%s'", key)
			}

			b.contents = append(b.contents, content{
				number:     num,
				contentBag: c,
			})
		}
	}

	return bags
}

func find(b, target *bag) bool {

	if b == target {
		// this doesn't count
		return false
	}

	//bfs
	queue := []*bag{b}

	for {
		if len(queue) == 0 {
			return false
		}

		curr := queue[0]
		queue = queue[1:]

		if curr == target {
			return true
		}

		for _, contents := range curr.contents {
			queue = append(queue, contents.contentBag)
		}
	}
}

func findChildBags(b *bag) int {
	if len(b.contents) == 0 {
		return 0
	}

	sum := 0
	for _, con := range b.contents {
		sum += con.number*findChildBags(con.contentBag) + con.number
	}

	return sum
}
