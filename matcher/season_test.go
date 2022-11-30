package matcher_test

import (
	"fmt"
	"testing"

	"github.com/gyu-young-park/lck_data_generator/matcher"
)

func TestIsMatch(t *testing.T) {
	seasonWithDateMatcher := matcher.NewLCKSeasonAndFandomSeasonMatcher()
	season := seasonWithDateMatcher.Match("2019-04-15")
	fmt.Println(season)
}
