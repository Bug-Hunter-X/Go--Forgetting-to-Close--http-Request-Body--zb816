# Go: Forgetting to Close `http.Request.Body`

This repository demonstrates a common but easily overlooked error in Go: forgetting to close the body of an HTTP request.  Failing to close the request body can lead to resource leaks, especially when handling a large number of requests.

The `bug.go` file shows the problematic code, while `bugSolution.go` provides the corrected version.

**Problem:**
The `handleRequest` function in `bug.go` does not explicitly close the request body. This can lead to open file descriptors, eventually exhausting system resources.

**Solution:**
The `bugSolution.go` file demonstrates the correct way to handle the request body: using `defer r.Body.Close()` ensures that the body is closed regardless of whether there are errors or panics during processing.