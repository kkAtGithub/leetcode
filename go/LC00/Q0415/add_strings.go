package Q0415

import "bytes"

const (
	nine uint8 = '9'
	one  uint8 = '1'
	zero uint8 = '0'
)

func addStrings(num1 string, num2 string) string {
	num1Digit := len(num1)
	num2Digit := len(num2)
	if num2Digit > num1Digit {
		return addStrings(num2, num1)
	}

	var rSumBuf bytes.Buffer
	rSumBuf.Grow(num1Digit + 1)

	var carry = zero
	var sum uint8
	i, j := num1Digit-1, num2Digit-1
	for i >= 0 && j >= 0 {
		carry, sum = add(num1[i], num2[j], carry)
		rSumBuf.WriteByte(sum)

		i--
		j--
	}

	for i >= 0 {
		carry, sum = add(num1[i], carry)
		rSumBuf.WriteByte(sum)
		i--
	}
	if carry == one {
		rSumBuf.WriteByte(carry)
	}

	var resDigit = rSumBuf.Len()
	var res = make([]byte, resDigit)
	for index, b := range rSumBuf.Bytes() {
		res[resDigit-index-1] = b
	}

	return string(res)
}

func add(num ...uint8) (uint8, uint8) {
	var sum = zero
	var carry uint8 = 0
	for _, n := range num {
		sum += n - zero
	}
	if sum > nine {
		sum -= 10
		carry = one
	} else {
		carry = zero
	}
	return carry, sum
}
