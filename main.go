// package main

// import (
// 	"fmt"
// 	back "groupie-tracker/backend"
// 	"log"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// const (
// 	artistsUrl  string = "https://groupietrackers.herokuapp.com/api/artists"
// 	locationUrl string = "https://groupietrackers.herokuapp.com/api/locations"
// 	dateUrl     string = "https://groupietrackers.herokuapp.com/api/dates"
// 	relationUrl string = "https://groupietrackers.herokuapp.com/api/relation"
// )

// var (
// 	serverRunning bool
// 	serverMutex   sync.Mutex
// )

// func startServer() {
// 	serverMutex.Lock()
// 	defer serverMutex.Unlock()

// 	if serverRunning {
// 		return
// 	}

// 	back.Client = &http.Client{Timeout: 10 * time.Second}
// 	//Synchronized goroutines
// 	wg := &sync.WaitGroup{}

// 	wg.Add(4)
// 	go func(wg *sync.WaitGroup) {
// 		back.GetJson(artistsUrl, &back.Myartists)
// 		wg.Done()
// 	}(wg)
// 	go func(wg *sync.WaitGroup) {
// 		back.GetJson(locationUrl, &back.Mylocations)
// 		wg.Done()
// 	}(wg)
// 	go func(wg *sync.WaitGroup) {
// 		back.GetJson(dateUrl, &back.Mydates)
// 		wg.Done()
// 	}(wg)
// 	go func(wg *sync.WaitGroup) {
// 		back.GetJson(relationUrl, &back.Myrelations)
// 		wg.Done()
// 	}(wg)
// 	wg.Wait()
// 	back.AppendToArtistStruct()

// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	http.HandleFunc("/", back.MainpageHandler)
// 	fmt.Println("💻 Server listening at")
// 	fmt.Println("👉 http://localhost:8084")
// 	fmt.Println("❌ CTRL+C to STOP Server")
// 	log.Fatal(http.ListenAndServe(":8084", nil))

// 	serverRunning = true
// }

// func stopServer() {
// 	serverMutex.Lock()
// 	defer serverMutex.Unlock()

// 	if !serverRunning {
// 		return
// 	}

// 	// Implement logic to stop the server here, if necessary

// 	serverRunning = false
// }

// func main() {
// 	http.HandleFunc("/start-server", func(w http.ResponseWriter, r *http.Request) {
// 		startServer()
// 		fmt.Fprintln(w, "Server started successfully")
// 	})

// 	http.HandleFunc("/stop-server", func(w http.ResponseWriter, r *http.Request) {
// 		stopServer()
// 		fmt.Fprintln(w, "Server stopped successfully")
// 	})

// 	fmt.Println("Server control endpoints:")
// 	fmt.Println("👉 http://localhost:8084/start-server (Start Server)")
// 	fmt.Println("👉 http://localhost:8084/stop-server (Stop Server)")

// 	select {}
// }

package main

import (
	"fmt"
	back "groupie-tracker/backend"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	artistsUrl  string = "https://groupietrackers.herokuapp.com/api/artists"
	locationUrl string = "https://groupietrackers.herokuapp.com/api/locations"
	dateUrl     string = "https://groupietrackers.herokuapp.com/api/dates"
	relationUrl string = "https://groupietrackers.herokuapp.com/api/relation"
)

func main() {
	back.Client = &http.Client{Timeout: 10 * time.Second}
	//Synchronized goroutines
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go func(wg *sync.WaitGroup) {
		back.GetJson(artistsUrl, &back.Myartists)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		back.GetJson(locationUrl, &back.Mylocations)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		back.GetJson(dateUrl, &back.Mydates)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		back.GetJson(relationUrl, &back.Myrelations)
		wg.Done()
	}(wg)
	wg.Wait()
	back.AppendToArtistStruct()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", back.MainpageHandler)
	fmt.Println("💻 Server listening at")
	fmt.Println("👉 http://localhost:8084")
	fmt.Println("❌ CTRL+C to STOP Server")
	log.Fatal(http.ListenAndServe(":8084", nil))

}
