# asciiweb-v2

# Terminal Test Script (test_endpoints.sh)

    This shell script acts as an automated collection of curl requests. You can run this directly from their terminal to verify if their 7 endpoints match your performance requirements.
    Create a file named test_endpoints.sh, paste the code below, and make it executable using chmod +x test_endpoints.sh.
    Use ./test_endpoints.sh  to run the tests.



# questin one

  #  Exercise 1: The Method Inspector
    
    Goal
    Build a /method-inspector endpoint that reads the HTTP method of every incoming request and echoes it back in a descriptive sentence. No method should be rejected — this handler accepts everything and reports what it sees.
    
        Key Tasks
        ●     Register a /method-inspector handler using http.HandleFunc.
        ●     Read the request method using r.Method.
        ●     Respond with a plain text message that includes the method name.
            ○     GET request → "You made a GET request."
            ○     POST request → "You made a POST request."
            ○     Any other method → "You made a [METHOD] request."
        ●     Do not hardcode each method with its own if/else branch — use the value of r.Method directly in your response string.

    Why this matters —
    In ascii-art-web your POST /ascii-art handler must distinguish incoming methods
    before doing any work. This exercise builds the muscle of reading r.Method
    cleanly and using it — not just checking it. 



# question two

  #  Exercise 2: The Echo Chamber
    
    Goal
    Create an /echo endpoint that only accepts POST requests. When a client sends a POST with a body, read the entire body and send it straight back. The response must be exactly what was sent — nothing added, nothing removed.
        
    Key Tasks
    ●     Reject any non-POST request with http.StatusMethodNotAllowed (405).
    ●     Read the full request body using io.ReadAll(r.Body).
    ●     Always defer r.Body.Close() immediately after checking for an error — not at the end of the function.
    ●     If the body is completely empty (zero bytes), return a 400 Bad Request with the message "body cannot be empty".
    ●     Write the body content back to w exactly as received.

        Think about —
        What does io.ReadAll return if the request has no body at all?
        Is len(body) == 0 the same as body == nil? Try both and see.
        What happens to r.Body if you read it twice without closing it? 

    Stretch — do this after the core task works
    ●     Add a response header: Content-Type: text/plain before writing the body back.
    ○     Use w.Header().Set("Content-Type", "text/plain") — and call it before w.Write().
    ○     What happens if you call w.Header().Set() after w.Write()? Try it and explain what you observe.





# question three

   # Exercise 3: Header Detective
    
    Goal
        Create a /headers endpoint that inspects two specific request headers: X-Custom-Token and Content-Type. The handler reads both, reports what it found, and enforces a rule about one of them.
        
    Key Tasks
        ●     Read X-Custom-Token using r.Header.Get("X-Custom-Token").
        ●     If X-Custom-Token is missing or empty — respond with 400 Bad Request and the message: "X-Custom-Token header is missing".
        ●     If X-Custom-Token is present — respond with a message that includes its value. Example: "Token received: abc123".
        ●     Also read Content-Type and append it to the response. If it is missing, append "Content-Type not provided".
        ●     The full response for a valid request must look like this:
            ○     Token received: abc123
            ○     Content-Type: application/json

    Why this matters —
        ascii-art-web reads r.Header indirectly through template and form handling.
        Understanding how headers work — and what happens when they are absent —
        prepares you for writing handlers that behave correctly under any input.

    Stretch — do this after the core task works
        ●     What does r.Header.Get() return for a header key that was never sent? Write a one-sentence answer in a comment at the top of your file.
        ●     Is r.Header.Get("x-custom-token") the same as r.Header.Get("X-Custom-Token")? Find out.



#  four

   # Exercise 4: Form Decoder
    Goal
        Build a /form endpoint that accepts a POST request with a URL-encoded form body containing two fields: username and language. Parse both, validate them, and return a formatted confirmation. This is the closest exercise in this set to what ascii-art-web actually does.
        
    Key Tasks
        ●     Reject non-POST requests with 405.
        ●     Call r.ParseForm() to parse the incoming form body — do not skip this step.
        ●     Read username and language using r.FormValue().
        ●     If either field is empty — return 400 Bad Request with a message identifying which field is missing.
            ○     Missing username → "username is required"
            ○     Missing language → "language is required"
        ●     If both are present — respond with: "Hello [username], you are coding in [language]!"

    
    Think about —
        What is the difference between r.ParseForm() + r.Form.Get() and just r.FormValue()?
        r.FormValue() calls ParseForm internally — but calling ParseForm explicitly first
        gives you control over error handling. Try it both ways.

    Stretch — do this after the core task works
            ●     Handle the case where the request Content-Type is not application/x-www-form-urlencoded. Return a 415 Unsupported Media Type with a clear message.
                ○     Read Content-Type from r.Header.Get() and check it before parsing.
                ○     Test it: curl -X POST -H "Content-Type: text/plain" -d "username=Ada" http://localhost:8080/form



#  five

# Exercise 5: Status Code Factory

    Goal
        Build a /status endpoint that accepts a code query parameter containing any HTTP status code number. The server must respond using that exact status code. This forces you to think about how status codes are set — and when they cannot be changed.
        
    Key Tasks
        ●     Read the code query parameter using r.URL.Query().Get("code").
        ●     If code is missing or empty — return 400 with the message: "code parameter is required".
        ●     Convert code to an integer using strconv.Atoi(). If conversion fails — return 400 with: "code must be a valid integer".
        ●     If the integer is not between 100 and 599 — return 400 with: "code must be a valid HTTP status code (100–599)".
        ●     If the code is valid — respond with that exact status code using w.WriteHeader(code) and a body message: "Responding with status [code]".
    
    Critical rule —
        You must call w.WriteHeader(code) BEFORE writing anything to w with w.Write()
        or fmt.Fprintf(). If you write the body first, Go automatically sends a 200 header
        and you cannot change it afterwards. The order is: WriteHeader → then Write.
        Test this deliberately: call w.Write() first, then w.WriteHeader(404). What does
        curl -v show you? Write your observation in a comment in your file.

    Stretch — do this after the core task works
        ●     After calling w.WriteHeader(), append a descriptive name to the body message.
            ○     ?code=404 → "Responding with status 404 Not Found"
            ○     Use http.StatusText(code) to get the official status name.



#  six

   # Exercise 6: The API Subtree

   **Goal**
        Build a mini API under the /api/v1/ path prefix using a separate http.ServeMux — not the default mux. This mux handles only /api/v1/ routes. Register two handlers inside it: /api/v1/ping and /api/v1/greet. Mount the whole submux onto the main server at /api/.
        
   **What a ServeMux subtree is**
        Go's http.ServeMux uses a simple rule: a pattern that ends in / matches any path that starts with it. This makes it possible to create a sub-router — a separate ServeMux that handles a subtree of routes — and mount it onto the main mux with http.StripPrefix. The main mux passes all /api/ requests to the submux, which handles them as if the /api prefix did not exist.
        
        // The pattern for a subtree — note the trailing slash
        mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
        
        // Inside apiMux, routes are registered WITHOUT /api
        apiMux.HandleFunc("/v1/ping", pingHandler)
        apiMux.HandleFunc("/v1/greet", greetHandler)
        
        // A request to /api/v1/ping:
        // mainMux strips "/api" → apiMux sees "/v1/ping" → routes to pingHandler

    
   **Key Tasks**
        ●     Create a new mux: apiMux := http.NewServeMux()
        ●     Register /v1/ping on apiMux — responds with "pong" in plain text.
        ●     Register /v1/greet on apiMux — reads a name query parameter and responds with "Greetings, [name]!" or "Greetings, Stranger!" if name is missing.
        ●     Mount apiMux onto the main mux:
            ○     mainMux := http.NewServeMux()
            ○     mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
        ●     Start the server using http.ListenAndServe(":8080", mainMux) — not nil.
        ●     Visiting /api/v1/ping must return "pong". Visiting /api/v1/greet?name=Zion must return "Greetings, Zion!".

   **Why this matters —**
        ascii-art-web mounts handlers at / and /ascii-art. A real application groups
        related routes under a prefix — /api/, /admin/, /v2/. ServeMux subtrees are
        Go's standard way to do this without an external router. Understanding
        StripPrefix is the foundation of any multi-route Go server.

   **Think about — write your answers in comments at the top of your file**
        ●     What happens if you use mainMux.Handle("/api", ...) without the trailing slash? Try it.
        ●     What does http.StripPrefix do — what would apiMux receive if you did NOT use StripPrefix?
        ●     What does it mean that http.ListenAndServe takes a Handler interface — and how does a *http.ServeMux satisfy that interface?


# seven


   **Exercise 7: Template Renderer**
            Goal
            Build a /render endpoint that accepts two query parameters — title and body — and renders them into an inline HTML template. The template must be defined inside your Go file as a string constant, parsed with html/template, and executed into the ResponseWriter. No external HTML files.
            
   **Key Tasks**
            ●     Define an HTML template as a raw string constant inside your .go file:
            
            const tmplStr = `
            <!DOCTYPE html>
            <html>
            <head><title>{{.Title}}</title></head>
            <body>
            <h1>{{.Title}}</h1>
            <p>{{.Body}}</p>
            </body>
            </html>
            `
            
            type PageData struct {
                Title string
                Body  string
            }

        
        ●     Parse the template using template.Must(template.New("page").Parse(tmplStr)).
        ●     In the /render handler — read title and body from the query string.
        ●     If either is missing — return 400 with: "title and body are required".
        ●     Execute the template with tmpl.Execute(w, PageData{Title: title, Body: body}).
        ●     If Execute returns an error — return 500 with: "template execution failed".
        ●     Set the Content-Type header to text/html before executing the template.
        
   **Critical difference from Exercise 2 —**

        In Exercise 2 you set Content-Type before writing the body. Here you must also
        set it before calling tmpl.Execute() — because Execute writes to w directly.
        Once Execute writes its first byte, headers are locked. Set them first.
        What is template.Must() — and when does it panic? Write the answer in a comment.

   **Stretch — do this after the core task works**

        ●     Add a third query parameter: style. If style=bold, wrap {{.Body}} in <strong> tags.
            ○     Hint: you cannot use an if statement in the query param — use template conditionals: {{if eq .Style \"bold\"}}<strong>{{.Body}}</strong>{{else}}{{.Body}}{{end}}
        ●     Add a PageData field Style string and pass it through from the query param.
