package main

import "testing"

func TestFilter(t *testing.T) {
	str := "aaa"
	forbidden := []string{"ab", "cd", "pq", "xy"}
	if !(containsVowels(str, 3) && hasLetterTwiceInARow(str) && !containsForbidden(str, forbidden)){
		t.FailNow()
	}
}