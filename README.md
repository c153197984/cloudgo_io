# CloudGo I/O

An static http server by Golang, which serves my personal blog as an example.

## Implementation

### Static File Server

All static files are under static/, including js, css, html and some footages.

With `PathPrefix` in `net/http`, we can easy serve those static files.

    // Handle static files(my blog)
    router.PathPrefix("/").Handler(
        http.StripPrefix("/", http.FileServer(http.Dir(webRoot+"./static/"))))

### Simple Javascript

The .js file `hello.js` sends a post request to server with `name` and `age` in
the page `/hello`, and server will reply a JSON object with fields `name` and
`age`.

We can see server output of the data.

    E.g.
    Hello from Alice [ 20 ]
    [negroni] 2017-11-19T17:53:07+08:00 | 200 |      460.827µs | localhost:3000 | POST /api/hello

### Form Submition

In page `/hello`, the website receives two values `name` and `age` and sends
them to server. Server will reply a JSON object with fields `name` and `age` and
the website will add the information to the table below.

    // Handle form submitted by user
    router.HandleFunc("/api/hello",
        func(res http.ResponseWriter, req *http.Request) {
            name := req.FormValue("name")
            age := req.FormValue("age")
            fmt.Println("Hello from", name, "[", age, "]")
            formatter := render.New(render.Options{
	           IndentJSON: true,
            })
            formatter.JSON(res, http.StatusOK, struct {
                Name string `json:"name"`
                Age  string `json:"age"`
            }{Name: name, Age: age})
        })

### StatusNotImplemented

If you visit `/unknown`, server will reply a `501 not implemented`, which is
implemented by a handler `NotImplementedHandler` by myself.

    func NotImplemented(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "501 page not implemented", http.StatusNotImplemented)
    }

    func NotImplementedHandler() http.Handler {
        return http.HandlerFunc(NotImplemented)
    }

    // In func main:
    // Handle unknown, send 501 not implemented
    router.Handle("/unknown", NotImplementedHandler())
