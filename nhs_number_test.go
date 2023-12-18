package nhs_number

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name      string
		nhsNumber string
		want      bool
	}{
		{
			name:      "valid NHS number",
			nhsNumber: "5990128088",
			want:      true,
		},
		{
			name:      "valid NHS number",
			nhsNumber: "1275988113",
			want:      true,
		},
		{
			name:      "valid NHS number",
			nhsNumber: "4536026665",
			want:      true,
		},
		{
			name:      "invalid NHS number",
			nhsNumber: "5990128087",
			want:      false,
		},
		{
			name:      "invalid NHS number",
			nhsNumber: "4536016660",
			want:      false,
		},
		{
			name:      "invalid NHS number",
			nhsNumber: "", // too short
			want:      false,
		},
		{
			name:      "invalid NHS number",
			nhsNumber: "123456789", // 1 digit missing
			want:      false,
		},
		{
			name:      "invalid NHS number",
			nhsNumber: "12345678901", // 1 digit too many
			want:      false,
		},
		{
			name:      "invalid NHS number",
			nhsNumber: "abcdefghij", // not a number - oddly adds up to 11 * n mod 11 = 0 if using runes
			want:      false,
		},
		{
			name:      "invalid NHS number",
			nhsNumber: "ABCDEFGHIJ", // not a number - same as above.
			want:      false,
		},
		{
			name:      "invalid NHS number",
			nhsNumber: "ABCDE12345", // half number, half not
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"/"+tt.nhsNumber, func(t *testing.T) {
			if got := IsValid(tt.nhsNumber); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

// bonus tests
func FuzzIsValid(f *testing.F) {
	f.Add("5990128088")
	f.Add("1275988113")
	f.Add("4536026665")
	f.Add("5990128087")
	f.Add("4536016660")
	f.Add("")
	f.Add("123456789")
	f.Add("12345678901")
	f.Add("abcdefghij")
	f.Add("ABCDEFGHIJ")
	f.Add("ABCDE12345")
	f.Fuzz(func(t *testing.T, nhsNumber string) {
		// ignore return value, just checking for panics
		_ = IsValid(nhsNumber)
	})
}
