package cmd

import (
	"fmt"
	"log"
	"math"
	"strings"

	"net/http"
	_ "net/http/pprof"

	"github.com/netr/aoc/util"
	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Day 5",
	Long:  `Day 5`,
	Run: func(cmd *cobra.Command, args []string) {
		lines := util.ReadFileIntoSlice("data/day5.txt")
		fmt.Println("URL: https://adventofcode.com/2023/day/5")

		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()

		fmt.Println("Part1:", getLowestSeedLocation(lines, false))
		fmt.Println("Part2:", getLowestSeedLocation(lines, true))
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}

type CategoryRange struct {
	Min int
	Max int
}

func newCategoryRange(min, maxRange string) CategoryRange {
	return CategoryRange{
		Min: util.MustAtoi(min),
		Max: util.MustAtoi(min) + util.MustAtoi(maxRange) - 1,
	}
}

type AlmanacCategory struct {
	Source      CategoryRange
	Destination CategoryRange
}

func (a AlmanacCategory) GetDest(source int) int {
	// check if source is in range
	if source < a.Source.Min || source > a.Source.Max {
		return 0
	}

	// return minimum destination with offset
	return a.Destination.Min + (source - a.Source.Min)
}

type AlmanacMap map[string][]AlmanacCategory

func (a AlmanacMap) GetDest(mapName string, source int) int {
	// iterate through all categories in map and check if source is in range
	// if source has a valid destination, return destination
	dest := 0
	for _, category := range a[mapName] {
		dest = category.GetDest(source)
		if dest != 0 {
			return dest
		}
	}
	// default to source
	return source
}

func createSeedRange(seeds []int) []int {
	total := 0
	// iterate seeds with a step of 2 and collect total seeds required
	for i := 0; i < len(seeds); i += 2 {
		curSeed := seeds[i]
		nextSeed := curSeed + seeds[i+1] - 1
		total += nextSeed - curSeed + 1
	}

	// preallocate newSeeds
	newSeeds := make([]int, total)
	idx := 0
	for i := 0; i < len(seeds); i += 2 {
		// iterate from curSeed to nextSeed and set newSeeds[idx]
		curSeed := seeds[i]
		nextSeed := curSeed + seeds[i+1] - 1
		for j := curSeed; j <= nextSeed; j++ {
			newSeeds[idx] = j
			idx++
		}
	}

	return newSeeds
}

func parseAlmanac(lines []string, seedsInRange bool) ([]int, AlmanacMap) {
	// parse seeds
	seedsSplit := strings.Split(lines[0], ": ")[1]
	seeds := util.SplitToInts(seedsSplit, " ")

	if seedsInRange {
		seeds = createSeedRange(seeds)
	}

	seedMap := make(AlmanacMap)
	buildingMap := false
	curMapName := ""

	for _, line := range lines[2:] {
		// check if a new map needs to be built and initialize it
		if strings.Contains(line, "map") {
			buildingMap = true
			curMapName = strings.Split(line, " ")[0]
			seedMap[curMapName] = []AlmanacCategory{}
		} else if buildingMap && line != "" {
			// build category. this will still work if there is no last line
			// since it's building it as it iterates through the lines
			nums := strings.Split(line, " ")
			seedMap[curMapName] = append(seedMap[curMapName], AlmanacCategory{
				Destination: newCategoryRange(nums[0], nums[2]),
				Source:      newCategoryRange(nums[1], nums[2]),
			})
		} else {
			// empty line, stop building current map
			buildingMap = false
		}
	}
	return seeds, seedMap
}

// grossly unoptimized. takes about 5 minutes to run
// the pprof output shows that the majority of the time is spent in map access
func getLowestSeedLocation(lines []string, seedsInRange bool) int {
	order := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}
	seeds, seedMap := parseAlmanac(lines, seedsInRange)

	lowest := math.MaxInt64
	carry := 0
	totalSeeds := len(seeds)

	for idx, seed := range seeds {
		carry = seed
		for _, mapName := range order {
			carry = seedMap.GetDest(mapName, carry)
		}
		if carry < lowest {
			lowest = carry
		}

		if idx%21000000 == 0 {
			pctLeft := 100 - float64(totalSeeds-idx)/float64(totalSeeds)*100
			log.Printf("idx: %d, lowest: %d, total: %d (%.2f%%)\n",
				idx,
				lowest,
				totalSeeds,
				pctLeft)
		}
	}

	return lowest
}
