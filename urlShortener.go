package main

import (
    "log"
    "net/http"
)

func redirectToMe(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "https://www.github.com/fluctix", 301)
}

func redirectToGo(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "https://golang.org/", 301)
}

func redirectToFB(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "https://www.facebook.com/", 301)
}
func redirectToGithub(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "https://www.github.com/fluctix", 301)
}

func redirectToTwitter(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "https://www.twitter.com/", 301)
}

func redirectToYoutube(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "https://www.youtube.com/", 301)
}

func redirectToGoogle(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "https://www.google.com/", 301)
}

func main() {
    http.HandleFunc("/me", redirectToMe)
   	http.HandleFunc("/go", redirectToGo)
   	http.HandleFunc("/fb", redirectToFB)
   	http.HandleFunc("/github", redirectToGithub)
   	http.HandleFunc("/tweet", redirectToTwitter)
   	http.HandleFunc("/youtube", redirectToYoutube)
   	http.HandleFunc("/google", redirectToGoogle)
   	err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    
}