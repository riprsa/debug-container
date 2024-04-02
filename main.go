package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Hi, {{ . }}!</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      background-color: #333;
      margin: 0;
      padding: 0;
      display: flex;
      height: 100vh;
      justify-content: center;
      align-items: center;
    }

    .container {
      text-align: center;
      padding: 20px;
      background-color: #444;
      border-radius: 8px;
      box-shadow: 0px 0px 10px rgba(255, 255, 255, 0.1);
    }

    h1 {
      color: #eee;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>Hello, {{ . }}!</h1>
  </div>
</body>
</html>
`

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":80"
	}

	containerID := os.Getenv("HOSTNAME")
	if containerID == "" {
		containerID = "World"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tmpl := template.Must(template.New("index").Parse(tmpl))
			err := tmpl.Execute(w, containerID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"container_id\":\"%s\"}", containerID)
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Start hw server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
