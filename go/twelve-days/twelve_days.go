package twelve

import (
	"fmt"
)

var numbers = []string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var gifts = []string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

func Verse(verseNum int) string {
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me", numbers[verseNum])
	and := "and "
	if verseNum == 1 {
		and = ""
	}
	for i := verseNum; i > 1; i-- {
		verse = fmt.Sprintf("%s, %s", verse, gifts[i])
	}
	verse = fmt.Sprintf("%s, %s", verse, and+gifts[1])
	return verse + "."
}

func Song() string {
	song := Verse(1) + "\n"
	for i := 2; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}
