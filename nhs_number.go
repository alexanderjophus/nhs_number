package nhs_number

func isValid(nhsNumber string) bool {
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
