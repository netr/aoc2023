package cmd

import "testing"

func Test_Day1(t *testing.T) {
	examples := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	expected := []int{12, 38, 15, 77}
	finalExpected := 142

	ttl := 0
	for i, example := range examples {
		result := calibrate(example)
		if result != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], result)
		}
		ttl += result
	}

	if ttl != finalExpected {
		t.Errorf("Expected %d, got %d", finalExpected, ttl)
	}
}

func Test_Day1_WordNums(t *testing.T) {
	examples := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := []int{29, 83, 13, 24, 42, 14, 76}
	finalExpected := 281

	ttl := 0
	for i, example := range examples {
		result := calibrateWordNums(example)
		if result != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], result)
		}
		ttl += result
	}

	if ttl != finalExpected {
		t.Errorf("Expected %d, got %d", finalExpected, ttl)
	}
}
