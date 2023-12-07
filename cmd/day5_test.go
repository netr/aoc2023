package cmd

import (
	"log"
	"testing"
	"time"
)

func Test_ParseAlmanac(t *testing.T) {
	expected := 4
	lines := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
	}

	seeds, seedMap := parseAlmanac(lines, false)
	if len(seeds) != expected {
		t.Errorf("seeds should be %d, got %d", expected, len(seeds))
	}
	if seedMap["seed-to-soil"] == nil {
		t.Errorf("seed-to-soil map should exist")
	}
}

func Test_AlmanacMap_GetDest(t *testing.T) {
	lines := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
	}

	table := []struct {
		seed int
		soil int
	}{
		{seed: 79, soil: 81},
		{seed: 14, soil: 14},
		{seed: 55, soil: 57},
		{seed: 13, soil: 13},
	}

	_, seedMap := parseAlmanac(lines, false)

	for _, row := range table {
		if seedMap.GetDest("seed-to-soil", row.seed) != row.soil {
			t.Errorf("seed-to-soil map should have dest of %d, got %d",
				row.soil, seedMap.GetDest("seed-to-soil", row.seed))
		}
	}
}

func Test_SolveDay5(t *testing.T) {
	lines := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}

	if getLowestSeedLocation(lines, false) != 35 {
		t.Errorf("should be 35, got %d", getLowestSeedLocation(lines, false))
	}

	if getLowestSeedLocation(lines, true) != 46 {
		t.Errorf("should be 46, got %d", getLowestSeedLocation(lines, true))
	}
}

func Test_CreateSeedRange(t *testing.T) {
	t.Skip() // takes too long to run on some machines
	seeds := []int{
		1310704671, 312415190, 1034820096, 106131293, 682397438, 30365957, 2858337556, 1183890307, 665754577, 13162298, 2687187253, 74991378, 1782124901, 3190497, 208902075, 226221606, 4116455504, 87808390, 2403629707, 66592398,
	}
	start := time.Now()
	seedRange := createSeedRange(seeds)
	elapsed := time.Since(start)

	log.Printf("createSeedRange took %s", elapsed)
	if len(seedRange) != 2104769314 {
		t.Errorf("seedRange should be 2104769314, got %d", len(seedRange))
	}
	if seedRange[3] != 1310704674 {
		t.Errorf("last seed should be 1310704674, got %d", seedRange[3])
	}
	if seedRange[312415190] != 1034820096 {
		t.Errorf("last seed should be 1034820096, got %d", seedRange[312415190])
	}
}
