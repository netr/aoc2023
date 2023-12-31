package cmd

import "testing"

func Test_TryCubeGame(t *testing.T) {
	game := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	if id, ok := tryCubeGame(game); !ok {
		t.Errorf("game %d should be possible", id)
	}
}

func Test_TryCubeGame_Larger(t *testing.T) {
	game := "Game 100: 3 blue, 3 red, 6 green; 7 red, 2 green, 16 blue; 14 green, 9 red, 9 blue; 8 red, 10 green, 9 blue; 6 blue, 11 red"
	if id, ok := tryCubeGame(game); ok {
		t.Errorf("game %d should not be possible", id)
	}
}

func Test_PlayCubeGames(t *testing.T) {
	games := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	count := playCubeGames(games)
	if count != 8 {
		t.Errorf("count should be 8, got %d", count)
	}
}

func Test_GetCubeGameMinimums(t *testing.T) {
	game := "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	power := getCubeGameMinimums(game)
	if power != 630 {
		t.Errorf("power should be 630, got %d", power)
	}
}

func Test_PlayCubeGamesForMinimums(t *testing.T) {
	games := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	count := playCubeGamesForMinimum(games)
	if count != 2286 {
		t.Errorf("count should be 8, got %d", count)
	}
}
