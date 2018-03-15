package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Match struct {
	team1, team2, result string
}

type ScoreCard map[string]*Score

type Score struct {
	wins, draws, losses int
}

func (s *Score) Points() int {
	return s.wins*3 + s.draws
}
func (s *Score) Matches() int {
	return s.wins + s.draws + s.losses
}

func (s ScoreCard) SortedKeys() []string {
	keys := make([]string, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}

	sort.Slice(keys, func(i, j int) bool {
		if s[keys[i]].Points() > s[keys[j]].Points() {
			return true
		} else if s[keys[i]].Points() == s[keys[j]].Points() && keys[i] < keys[j] {
			return true
		}
		return false
	})

	return keys
}

func Tally(r io.Reader, w io.Writer) error {
	s := make(ScoreCard)
	b, _ := ioutil.ReadAll(r)
	w.Write([]byte(fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")))
	for _, line := range strings.Split(string(b), "\n") {
		if line == "" || line[0] == '#' {
			continue
		}
		if strings.Count(line, ";") != 2 {
			return fmt.Errorf("Invalid Input")
		}

		parsed := strings.Split(line, ";")
		team1, team2, result := parsed[0], parsed[1], parsed[2]

		if s[team1] == nil {
			s[team1] = new(Score)
		}
		if s[team2] == nil {
			s[team2] = new(Score)
		}

		if result == "win" {
			s[team1].wins++
			s[team2].losses++
		} else if result == "draw" {
			s[team1].draws++
			s[team2].draws++
		} else if result == "loss" {
			s[team1].losses++
			s[team2].wins++
		} else {
			return fmt.Errorf("Invalid result")
		}
	}

	for _, team := range s.SortedKeys() {
		w.Write([]byte(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", team, s[team].Matches(), s[team].wins, s[team].draws, s[team].losses, s[team].Points())))
	}
	return nil
}
