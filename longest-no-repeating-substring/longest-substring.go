package longestNoRepeatingSubstring

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return len(s)
	}

	// init
	maxLength := 1
	left := 0
	right := 1
	currentLength := 1

	cMap := map[uint8]int{s[0]: 0}

	for right < len(s) {
		if left == right {
			right++
			continue
		}

		println("before", "left=", left, "right=", right, "currentLength=", currentLength, "maxLength=", maxLength)

		if v, exist := cMap[s[right]]; exist {
			// count length
			if maxLength < currentLength {
				maxLength = currentLength
			}

			for i := left; i < v+1; i++ {
				delete(cMap, s[i])
			}
			left = v + 1
			currentLength = 1
			continue
		}

		// if not
		cMap[s[right]] = right
		currentLength++
		right++
		println("after", "left=", left, "right=", right, "currentLength=", currentLength, "maxLength=", maxLength)

	}
	println("currentLength=", currentLength)

	println("maxLength=", maxLength)
	if maxLength < currentLength {
		maxLength = currentLength
	}

	println("maxLength=", maxLength)

	return maxLength
}

func Run() {

	str := "bbbbb"
	println(lengthOfLongestSubstring(str))

}
