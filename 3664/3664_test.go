package main

import "testing"

type ScoreTestInput struct {
	cards []string
	x     byte
}

func ScoreTest(t *testing.T, input ScoreTestInput, expect int) {
	result := score(input.cards, input.x)
	if result != expect {
		t.Errorf("Test case for %s failed expected %d but got %d", input.cards, expect, result)
	}
}

func TestScoreCase1(t *testing.T) {
	input := ScoreTestInput{
		cards: []string{"aa", "ab", "ba", "ac"},
		x:     'a',
	}
	ScoreTest(t, input, 2)
}

func TestScoreCase2(t *testing.T) {
	input := ScoreTestInput{
		cards: []string{"aa", "ab", "ba", "ac"},
		x:     'b',
	}
	ScoreTest(t, input, 0)
}

func TestScoreCase3(t *testing.T) {
	input := ScoreTestInput{
		cards: []string{"aa", "ab", "ba"},
		x:     'a',
	}
	ScoreTest(t, input, 1)
}

func TestScoreCase4(t *testing.T) {
	input := ScoreTestInput{
		cards: []string{"bb", "bb"},
		x:     'b',
	}
	ScoreTest(t, input, 0)
}

func TestScoreCase5(t *testing.T) {
	input := ScoreTestInput{
		cards: []string{"ba", "ba"},
		x:     'b',
	}
	ScoreTest(t, input, 0)
}

func TestScoreCase6(t *testing.T) {
	input := ScoreTestInput{
		cards: []string{"ba", "bc", "bc", "bb", "bb", "bb"},
		x:     'b',
	}
	ScoreTest(t, input, 3)
}

func TestScoreCase7(t *testing.T) {
	input := ScoreTestInput{
		cards: []string{"ax", "ay", "az", "aa", "aa", "ka", "ja"},
		x:     'a',
	}
	ScoreTest(t, input, 3)
}
