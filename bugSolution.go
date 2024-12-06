func handleRequest(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    // ... process the request body ...
}