package highscores
import "slices"
type HighScores struct{
    score []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
    
	return &HighScores{
		score: scores,
	}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.score
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
    
    return s.score[len(s.score)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
    max:= s.score[0]
	for _ , value := range s.score {
        if max< value{
            max = value
        }
    }
    return max
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
    slice := slices.Clone(s.score)
	slices.SortFunc(slice , func(a,b int ) int {
        return b-a
    })
    if len(slice)>=3{
        slice= slice[:3]
    }
    
    return slice
}
