package nhs_number

// IsValid returns true if the NHS number is valid, false otherwise.
//
// The NHS number is a 10 digit number.
// The first digit is multiplied by 10, the second by 9, the third by 8, etc.
// The sum of these multiplications is then divided by 11.
// If the remainder is 0, the NHS number is valid.
func IsValid(nhsNumber string) bool {
	if len(nhsNumber) != 10 {
		return false
	}

	weight := 10
	sum := 0
	for _, digit := range nhsNumber {
		if digit < '0' || digit > '9' {
			return false
		}
		sum += int(digit-'0') * weight
		weight--
	}

	return sum%11 == 0
}
