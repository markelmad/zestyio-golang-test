package helper

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	data := []ResultsFromLebenshtein{
		{Title: "All's Well That Ends Well", ReadCount: 0, Distance: 2},
		{Title: "Comedy of Errors", ReadCount: 0, Distance: 1},
		{Title: "Tempest", ReadCount: 0, Distance: 3},
		{Title: "Winter's Tale", ReadCount: 0, Distance: 2},
		{Title: "King John", ReadCount: 0, Distance: 4},
		{Title: "Timon of Athens", ReadCount: 0, Distance: 2},
	}
	t.Run("sort", func(t *testing.T) {
		Sort(data)

		for _, d := range data {
			fmt.Printf("Distance: %d, Title: %s, ReadCount: %d\n", d.Distance, d.Title, d.ReadCount)
		}

	})
}
