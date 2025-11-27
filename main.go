package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running at http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Go + HTMX + CSS</title>

	<!-- HTMX -->
	<script src="https://unpkg.com/htmx.org@1.9.10"></script>

	<style>
		body {
			font-family: sans-serif;
			background: #f4f4f4;
			margin: 0;
			padding: 40px;
			text-align: center;
		}
		.card {
			background: white;
			padding: 20px;
			margin: auto;
			width: 300px;
			border-radius: 8px;
			box-shadow: 0 2px 5px rgba(0,0,0,0.1);
		}
		button {
			background: #4CAF50;
			color: white;
			border: none;
			padding: 10px 20px;
			border-radius: 4px;
			cursor: pointer;
			font-size: 16px;
		}
		button:hover {
			background: #45a049;
		}
	</style>
</head>
<body>

<div class="card">
	<h2>HTMX + Go</h2>

	<div id="result"></div>

	<button 
		hx-get="/hello"
		hx-trigger="click"
		hx-target="#result"
		hx-swap="innerHTML"
	>
		Load Message
	</button>
</div>

</body>
</html>`)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<p style="color: #4CAF50; font-weight: bold;">
	Hello from the Go server! 🎉
</p>`)
}
