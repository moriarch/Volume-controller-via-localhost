package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itchyny/volume-go"
	"github.com/rs/cors"
)

func handleMain(w http.ResponseWriter, r *http.Request) {

}

func handleVolume(w http.ResponseWriter, r *http.Request) {
	vol, err := volume.GetVolume()
	if err != nil {
		print("err")
	}
	fmt.Printf("current volume: %d\n", vol)
	jsonBytes, err := json.Marshal(vol)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	w.Write(jsonBytes)

}

type Data struct {
	Volume int `json:"volume"`
}

func handleSetVolume(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, data.Volume)
	err = volume.SetVolume(data.Volume)
	if err != nil {
		print("err set volume")
	}
	fmt.Printf("set volume success\n")
}
func handleMute(w http.ResponseWriter, r *http.Request) {
	err := volume.Mute()
	if err != nil {
		print("error mute")
	}
	fmt.Printf("set mute success\n")
}
func handleUnmute(w http.ResponseWriter, r *http.Request) {
	err := volume.Unmute()
	if err != nil {
		print("error mute")
	}
	fmt.Printf("set unmute success\n")
}

func handlePing(w http.ResponseWriter, r *http.Request) {

	jsonBytes, err := json.Marshal(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if err != nil {
		print("err")
	}
	w.Write(jsonBytes)
}
func main() {

	mux := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./frontend/dist")).ServeHTTP(w, r)
	})
	mux.HandleFunc("/getVolume", handleVolume)
	mux.HandleFunc("/setVolume", handleSetVolume)
	mux.HandleFunc("/mute", handleMute)
	mux.HandleFunc("/unmute", handleUnmute)
	mux.HandleFunc("/ping", handlePing)

	handler := cors.Handler(mux)
	http.ListenAndServe(":80", handler)
}
