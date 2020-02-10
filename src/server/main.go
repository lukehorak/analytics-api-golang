package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type counters struct {
	// Counters no longer need sync.Mutex, as they're stored in
	// a syncMap which handles concurrent read/writes
	view  int
	click int
}

var (
	list    sync.Map
	content = []string{"sports", "entertainment", "business", "education"}
)

/*///////////////////////////////////////////
	makeKeyString()
		description:
			Takes the data type (from the content array) and returns a
			formatted date string to be used as a key in the list syncMap
		parameters:
			[data (string)]:
				Data type (e.g. sports, business), from the content array
		return value:
			Formatted string to be used as a key for the syncMap
///////////////////////////////////////////*/
func makeKeyString(data string) string {
	t := time.Now()
	tString := t.Format("2006-01-02 15:04")
	keyString := fmt.Sprintf("%s:%s", data, tString)
	return keyString
}

/*///////////////////////////////////////////
	processClick()
		description:
			Processes/records click actions taken by the "user", writing to list
		parameters:
			[key (string)]:
				Pre-formatted string (generated by makeKeyString() function) used as key for
				loading/storing in syncMap
		return value:
			nil if successful. error if failed
///////////////////////////////////////////*/
func processClick(key string) error {
	c, ok := list.Load(key)
	if ok {
		counter := c.(counters)
		counter.click++
		list.Store(key, counter)
	}
	return nil
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to EQ Works 😎")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	// Build keyString for syncMap
	data := content[rand.Intn(len(content))]
	key := makeKeyString(data)

	// Load counters for key, or init if nonexistent
	c, ok := list.LoadOrStore(key, counters{view: 1})
	counter := c.(counters)
	if ok {
		counter.view++
		list.Store(key, counter)
	}

	// Request error handling
	err := processRequest(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	// simulate random click call
	if rand.Intn(100) < 50 {
		processClick(key)
	}

	// returning string with response, for testing
	// TODO - remove this before submitting!
	returnString := fmt.Sprintf("%s --> %v", key, counter)
	fmt.Fprint(w, returnString)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	if !isAllowed() {
		w.WriteHeader(429)
		return
	}

	m := map[string]interface{}{}
	list.Range(func(key, value interface{}) bool {
		m[fmt.Sprint(key)] = fmt.Sprintf("%v", value)
		return true
	})

	b, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func processRequest(r *http.Request) error {
	time.Sleep(time.Duration(rand.Int31n(50)) * time.Millisecond)
	return nil
}

func isAllowed() bool {
	return true
}

func uploadCounters() error {
	return nil
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", welcomeHandler)
	mux.HandleFunc("/view/", viewHandler)
	mux.HandleFunc("/stats/", statsHandler)

	log.Fatal(http.ListenAndServe(":8080", rateLimit(mux)))
}
