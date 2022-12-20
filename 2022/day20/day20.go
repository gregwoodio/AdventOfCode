package day20

type node struct {
	val        int64
	prev, next *node
}

func solvePartOne(input []int) int64 {
	return solve(input, false)
}

func solvePartTwo(input []int) int64 {
	return solve(input, true)
}

func solve(input []int, isPartTwo bool) int64 {
	list := makeList(input, isPartTwo)
	head := list[0]

	cycles := 1
	if isPartTwo {
		cycles = 10
	}

	for cycle := 0; cycle < cycles; cycle++ {
		for _, n := range list {
			tmp := n
			if n.val == 0 {
				continue
			}

			// remove from list
			tmp.prev.next = tmp.next
			tmp.next.prev = tmp.prev

			val := n.val
			// for ; val < 0; val += len(list) {
			// }
			for ; val > int64(len(list)-1); val %= int64(len(list) - 1) {
			}

			if val < 0 {
				for i := val; i < 0; i++ {
					tmp = tmp.prev
				}
				tmp.prev.next = n
				n.prev = tmp.prev
				n.next = tmp
				tmp.prev = n
			} else if val > 0 {
				for i := int64(0); i < val; i++ {
					tmp = tmp.next
				}
				tmp.next.prev = n
				n.next = tmp.next
				n.prev = tmp
				tmp.next = n
			}

			// fmt.Printf("moved %d between %d and %d\n", n.val, n.prev.val, n.next.val)
		}
	}

	for ; head.val != 0; head = head.next {
	}

	for i := 0; i < 1000; i++ {
		head = head.next
	}
	var thousand int64 = head.val
	for i := 0; i < 1000; i++ {
		head = head.next
	}
	var twothousand int64 = head.val
	for i := 0; i < 1000; i++ {
		head = head.next
	}
	var threethousand int64 = head.val

	return thousand + twothousand + threethousand
}

func makeList(input []int, isPartTwo bool) []*node {
	list := []*node{}
	var prev *node
	for _, val := range input {
		var val64 int64 = int64(val)
		if isPartTwo {
			val64 *= 811589153
		}

		n := &node{
			val: val64,
		}
		if prev != nil {
			prev.next = n
			n.prev = prev
		}
		list = append(list, n)
		prev = n
	}

	// connect tail to head
	prev.next = list[0]
	list[0].prev = prev

	return list
}
