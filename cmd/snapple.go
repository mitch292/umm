package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// snappleCmd represents the snapple command
var snappleCmd = &cobra.Command{
	Use:   "snapple",
	Short: "Generate a random fact",
	Long: `
		Ever get bored and just want to hear a random fact? Same.
		Beware these are snapple facts...Who knows if they're true.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("snapple called")

		rand.Seed(time.Now().Unix())
		n := rand.Int() % len(snappleFacts)
		fmt.Println(snappleFacts[n].text)
	},
}

func init() {
	rootCmd.AddCommand(snappleCmd)
}

type snappleFactType string

const (
	sports        = "Sports"
	politics      = "Politics"
	animalKingdom = "Animal Kingdom"
	geography     = "Geography"
	random        = "Random"
)

type snappleFact struct {
	id       int
	category snappleFactType
	text     string
}

var snappleFacts = []snappleFact{
	snappleFact{
		id:       4,
		category: animalKingdom,
		text:     "Slugs have four noses.",
	},
	snappleFact{
		id:       8,
		category: animalKingdom,
		text:     "A bee has five eyelids.",
	},
	snappleFact{
		id:       9,
		category: animalKingdom,
		text:     "The average speed of a housefly is 4.5 mph.",
	},
	snappleFact{
		id:       38,
		category: animalKingdom,
		text:     "Fish cough.",
	},
	snappleFact{
		id:       61,
		category: animalKingdom,
		text:     "Pigs can get sunburn.",
	},
	snappleFact{
		id:       68,
		category: random,
		text:     `The longest one-syllable word is, "screeched"`,
	},
	snappleFact{
		id:       71,
		category: geography,
		text:     `There is a town called "Big Ugly" in West Virginia`,
	},
}
