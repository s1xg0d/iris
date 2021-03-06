package main

import (
	"encoding/xml"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// User example struct for json and msgpack.
type User struct {
	Firstname string `json:"firstname" msgpack:"firstname"`
	Lastname  string `json:"lastname" msgpack:"lastname"`
	City      string `json:"city" msgpack:"city"`
	Age       int    `json:"age" msgpack:"age"`
}

// ExampleXML just a test struct to view represents xml content-type
type ExampleXML struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

// ExampleYAML just a test struct to write yaml to the client.
type ExampleYAML struct {
	Name       string `yaml:"name"`
	ServerAddr string `yaml:"ServerAddr"`
}

func main() {
	app := iris.New()

	// Read
	app.Post("/decode", func(ctx iris.Context) {
		// Read https://github.com/kataras/iris/blob/master/_examples/request-body/read-json/main.go as well.
		var user User
		ctx.ReadJSON(&user)

		ctx.Writef("%s %s is %d years old and comes from %s!", user.Firstname, user.Lastname, user.Age, user.City)
	})

	// Write
	app.Get("/encode", func(ctx iris.Context) {
		u := User{
			Firstname: "John",
			Lastname:  "Doe",
			City:      "Neither FBI knows!!!",
			Age:       25,
		}

		// Manually setting a content type: ctx.ContentType("text/javascript")
		ctx.JSON(u)
	})

	// Other content types,

	app.Get("/binary", func(ctx iris.Context) {
		// useful when you want force-download of contents of raw bytes form.
		ctx.Binary([]byte("Some binary data here."))
	})

	app.Get("/text", func(ctx iris.Context) {
		ctx.Text("Plain text here")
	})

	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(map[string]string{"hello": "json"}) // or myjsonStruct{hello:"json}
	})

	app.Get("/jsonp", func(ctx iris.Context) {
		ctx.JSONP(map[string]string{"hello": "jsonp"}, context.JSONP{Callback: "callbackName"})
	})

	app.Get("/xml", func(ctx iris.Context) {
		ctx.XML(ExampleXML{One: "hello", Two: "xml"})
		// OR:
		// ctx.XML(iris.XMLMap("keys", iris.Map{"key": "value"}))
	})

	app.Get("/markdown", func(ctx iris.Context) {
		ctx.Markdown([]byte("# Hello Dynamic Markdown -- iris"))
	})

	app.Get("/yaml", func(ctx iris.Context) {
		ctx.YAML(ExampleYAML{Name: "Iris", ServerAddr: "localhost:8080"})
		// OR:
		// ctx.YAML(iris.Map{"name": "Iris", "serverAddr": "localhost:8080"})
	})

	// app.Get("/protobuf", func(ctx iris.Context) {
	// 	ctx.Protobuf(proto.Message)
	// })

	app.Get("/msgpack", func(ctx iris.Context) {
		u := User{
			Firstname: "John",
			Lastname:  "Doe",
			City:      "Neither FBI knows!!!",
			Age:       25,
		}

		ctx.MsgPack(u)
	})

	// http://localhost:8080/decode
	// http://localhost:8080/encode
	//
	// http://localhost:8080/binary
	// http://localhost:8080/text
	// http://localhost:8080/json
	// http://localhost:8080/jsonp
	// http://localhost:8080/xml
	// http://localhost:8080/markdown
	// http://localhost:8080/msgpack
	//
	// `iris.WithOptimizations` is an optional configurator,
	// if passed to the `Run` then it will ensure that the application
	// response to the client as fast as possible.
	//
	//
	// `iris.WithoutServerError` is an optional configurator,
	// if passed to the `Run` then it will not print its passed error as an actual server error.
	app.Listen(":8080", iris.WithOptimizations)
}
