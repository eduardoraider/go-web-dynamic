# Go Web Dynamic Page

This project serves a dynamic HTML page using Go. It utilizes the `net/http` package to create an HTTP server that renders HTML templates.

## Usage

To run the server, use the following command:

```
go run main.go
```

You can then access the dynamic HTML page at:

```
http://localhost:8080
```

## Implementation Details

- The `html/template` package is used to parse and execute HTML templates.
- Templates are located in the "templates" folder.
- The template is compiled using `template.Must(template.ParseFiles("templates/index.html"))`, which parses the template file and panics if there is an error during parsing.
- A handler function is defined to serve the template. Inside the handler function:
    - A `Data` struct with the title is created.
    - The template is executed with the data using `tmpl.Execute(w, data)`, where `tmpl` is the compiled template.
- The server is started on port 8080 using `http.ListenAndServe(":8080", nil)`.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer