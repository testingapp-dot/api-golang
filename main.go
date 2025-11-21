package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// Root page (HTML + CSS + HTMX)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Go + HTMX Demo</title>

	<!-- HTMX CDN -->
	<script src="https://unpkg.com/htmx.org@1.9.10"></script>

	<style>
		body {
			font-family: sans-serif;
			margin: 2rem;
			background: #f1f5f9;
		}
		.card {
			background: white;
			padding: 1rem 1.5rem;
			border-radius: 8px;
			box-shadow: 0 2px 8px rgba(0,0,0,0.1);
			max-width: 400px;
		}
		button {
			background: #2563eb;
			color: white;
			border: none;
			padding: 0.5rem 1rem;
			border-radius: 5px;
			cursor: pointer;
			margin-top: 1rem;
		}
		button:hover {
			background: #1e40af;
		}
	</style>
</head>
<body>

	<div class="card">
		<h2>Go + HTMX</h2>

		<div id="result" style="margin:1rem 0; color:#333;">
			Click the button to load content.
		</div>

		<button
			hx-get="/load"
			hx-target="#result"
			hx-swap="innerHTML"
		>
			Load Message
		</button>
	</div>

</body>
</html>
		`)
	})

	// HTMX endpoint (returns the current server time)
	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Format("15:04:05")
		fmt.Fprintf(w, "Hello from Go! The current server time is <b>%s</b>", now)
	})

	fmt.Println("Server running at http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}