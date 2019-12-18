package main

import (
	"strings"
	"testing"
)

func getTestbooks() books {
	return books{
		book{
			title:       "Ich helfe dir kochen. Das erfolgreiche Universalkochbuch mit großem Backteil",
			isbn:        "5554-5545-4518",
			authors:     "null-walter@echocat.org",
			description: "Auf der Suche nach einem Basiskochbuch steht man heutzutage vor einer Fülle von Alternativen. Es fällt schwer, daraus die für sich passende Mixtur aus Grundlagenwerk und Rezeptesammlung zu finden. Man sollte sich darüber im Klaren sein, welchen Schwerpunkt man setzen möchte oder von welchen Koch- und Backkenntnissen man bereits ausgehen kann.",
		},
		book{
			title:       "Genial italienisch",
			isbn:        "1024-5245-8584",
			authors:     "null-lieblich@echocat.org,null-walter@echocat.org,null-rabe@echocat.org",
			description: "Starkoch Jamie Oliver war mit seinem VW-Bus in Italien unterwegs -- und hat allerlei kulinarische Souvenirs mitgebracht. Es lohnt sich, einen Blick in sein Gepäck zu werfen...",
		},
	}

}

func TestBookToString(t *testing.T) {
	testbook := getTestbooks()[0]
	expected := "Ich helfe dir kochen. Das erfolgreiche Universalkochbuch mit großem Backteil[5554-5545-4518]\nnull-walter@echocat.org\nAuf der Suche nach einem Basiskochbuch steht man heutzutage vor einer Fülle von Alternativen. Es fällt schwer, daraus die für sich passende Mixtur aus Grundlagenwerk und Rezeptesammlung zu finden. Man sollte sich darüber im Klaren sein, welchen Schwerpunkt man setzen möchte oder von welchen Koch- und Backkenntnissen man bereits ausgehen kann."
	actual := testbook.toString()
	assert(t, expected, actual)
}

func TestBooksToString(t *testing.T) {
	testbook := getTestbooks()
	expected := "Ich helfe dir kochen. Das erfolgreiche Universalkochbuch mit großem Backteil[5554-5545-4518]\nnull-walter@echocat.org\nAuf der Suche nach einem Basiskochbuch steht man heutzutage vor einer Fülle von Alternativen. Es fällt schwer, daraus die für sich passende Mixtur aus Grundlagenwerk und Rezeptesammlung zu finden. Man sollte sich darüber im Klaren sein, welchen Schwerpunkt man setzen möchte oder von welchen Koch- und Backkenntnissen man bereits ausgehen kann.\n" +
		"Genial italienisch[1024-5245-8584]\nnull-lieblich@echocat.org,null-walter@echocat.org,null-rabe@echocat.org\nStarkoch Jamie Oliver war mit seinem VW-Bus in Italien unterwegs -- und hat allerlei kulinarische Souvenirs mitgebracht. Es lohnt sich, einen Blick in sein Gepäck zu werfen..."
	actual := testbook.toString()
	assert(t, expected, actual)
}

func TestNewbook(t *testing.T) {
	line := "Ich helfe dir kochen. Das erfolgreiche Universalkochbuch mit großem Backteil;5554-5545-4518;null-walter@echocat.org;Auf der Suche nach einem Basiskochbuch steht man heutzutage vor einer Fülle von Alternativen. Es fällt schwer, daraus die für sich passende Mixtur aus Grundlagenwerk und Rezeptesammlung zu finden. Man sollte sich darüber im Klaren sein, welchen Schwerpunkt man setzen möchte oder von welchen Koch- und Backkenntnissen man bereits ausgehen kann."
	expected := getTestbooks()[0]
	actual := newbook(line)
	assert(t, expected.toString(), actual.toString())
	actual = newbook("Ich helfe dir kochen. Das erfolgreiche Universalkochbuch mit großem Backteil;5554-5545-4518;null-walter@echocat.org")
	assert(t, "", actual.toString())
}

func TestLoad(t *testing.T) {
	file := "resources/books.csv"
	actual := loadBooks(file)
	assert(t, "Ich helfe dir kochen. Das erfolgreiche Universalkochbuch mit großem Backteil[5554-5545-4518]\nnull-walter@echocat.org\nAuf der Suche nach einem Basiskochbuch steht man heutzutage vor einer Fülle von Alternativen. Es fällt schwer, daraus die für sich passende Mixtur aus Grundlagenwerk und Rezeptesammlung zu finden. Man sollte sich darüber im Klaren sein, welchen Schwerpunkt man setzen möchte oder von welchen Koch- und Backkenntnissen man bereits ausgehen kann.", actual[0].toString())
	assert(t, "Schuhbecks Kochschule. Kochen lernen mit Alfons Schuhbeck[1215-4545-5895]\nnull-walter@echocat.org\nAller Anfang ist leicht! Zumindest, wenn man beim Kochenlernen einen Lehrer wie Alfons Schuhbeck zur Seite hat. Mit seiner Hilfe kann auch der ungeschickteste Anfänger beste Noten für seine Gerichte bekommen. Der Trick, den der vielfach ausgezeichnete Meisterkoch dabei anwendet, lautet visualisieren. Die einzelnen Arbeitsschritte werden auf Farbfotos im Format von ca. 3x4 cm abgebildet. Unter diesen stehen kurz und knapp die Angaben zur Zubereitung. So präsentiert Schuhbecks Kochschule alles bequem auf einen Blick. Und der interessierte Kochneuling kann sich auf die Hauptsache konzentrieren – aufs Kochen. Wie die Speise aussehen soll, zeigt jeweils das Farbfoto auf der linken Seite, auf der auch die Zutaten – dank des geschickten Layouts – ebenfalls sehr übersichtlich aufgelistet sind. Außerdem gibt Schuhbeck prägnante Empfehlungen zu Zutaten und Zubereitung.", actual[len(actual)-1].toString())
}

func TestPrintBooks(t *testing.T) {
	res := captureOutput(func() {
		getTestbooks().printBooks()
	})
	assertTrue(t, strings.Contains(res, "Ich helfe dir kochen"))
	assertTrue(t, strings.Contains(res, "Starkoch Jamie Oliver"))
}

func TestFindByISBN(t *testing.T) {
	books := getTestbooks()
	found := books.findByISBN("5554-5545-4518")
	expected := "Ich helfe dir kochen. Das erfolgreiche Universalkochbuch mit großem Backteil[5554-5545-4518]\nnull-walter@echocat.org\nAuf der Suche nach einem Basiskochbuch steht man heutzutage vor einer Fülle von Alternativen. Es fällt schwer, daraus die für sich passende Mixtur aus Grundlagenwerk und Rezeptesammlung zu finden. Man sollte sich darüber im Klaren sein, welchen Schwerpunkt man setzen möchte oder von welchen Koch- und Backkenntnissen man bereits ausgehen kann."
	assert(t, expected, found.toString())
	found = books.findByISBN("0000-0000-0000")
	expected = ""
	assert(t, expected, found.toString())
}
