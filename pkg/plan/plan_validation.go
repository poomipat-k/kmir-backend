package plan

import "fmt"

func validateScore(scores map[string]int) (string, error) {
	for i := 1; i <= 7; i++ {
		key := fmt.Sprintf("q_%d", i)
		val, exist := scores[key]
		if !exist {
			return key, ScoreRequiredError{}
		}
		if val < 1 || val > 10 {
			return key, ScoreValueOutOfRangeError{}
		}
	}

	return "", nil
}
