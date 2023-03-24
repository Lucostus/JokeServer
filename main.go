package JokeServer

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type Joke struct {
	gorm.Model
	Text string
}

var db = &gorm.DB{}
var brng = badRandomNumberGenerator()

func badRandomNumberGenerator() func() int {
	i := 0
	return func() int {
		i++
		i := i<<3 + 1
		return i % 9
	}
}

func GetRandomJoke() Joke {
	joke := Joke{}
	id := brng()
	fmt.Println("id", id)
	db.First(&joke, id)
	return joke
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	_, err := io.WriteString(w, fmt.Sprintf("Joke of the Day: %s", GetRandomJoke().Text))
	if err != nil {
		return
	}
}

func init() {
	err := error(nil)
	db, err = gorm.Open(sqlite.Open("jokes.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Joke{})
	if err != nil {
		return
	}

	var jokes = []Joke{
		{Text: "Why did the JavaScript developer wear glasses? Because he couldn't C#."},
		{Text: "Why did the database admin walk into a NoSQL bar? Because they wouldn't let him join."},
		{Text: "Why did the programmer go broke? He used up all his cache."},
		{Text: "Why was the database administrator so bad at telling jokes? He kept forgetting to join the punchline."},
		{Text: "Why did the programmer quit his job at the calendar factory? He took a day off."},
		{Text: "Why do programmers prefer dark mode? Less light means less bugs."},
		{Text: "Why did the programmer get stuck in the shower? He forgot to ESC."},
		{Text: "Why did the programmer get lost in the forest? He couldn't find his way through the trees."},
		{Text: "Why did the programmer's girlfriend leave him? He kept comparing her to an API."},
		{Text: "Why was the programmer always cold? He left his Windows open."},
		{Text: "Why did the programmer get lost in the forest? He couldn't find his way through the trees."},
	}
	db.Create(&jokes)
}
func StartHAHAHAHAServer() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("ListenAndServe failed: %v", err)
	}
}
