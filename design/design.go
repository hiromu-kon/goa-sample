package design

import dsl "goa.design/goa/v3/dsl"

var _ = dsl.API("todo", func() {
	dsl.Title("Todo Service")
	dsl.Description("Service for managing todo items")
	dsl.Version("0.1")
	dsl.Server("todo", func() {
		dsl.Host("localhost", func() { dsl.URI("http://localhost:8088") })
	})
})

var Todo = dsl.Type("Todo", func() {
	dsl.Description("Todo describes a todo item.")
	dsl.Attribute("id", dsl.String, "ID of the todo", func() {})
	dsl.Attribute("title", dsl.String, "Title of the todo", func() {})
	dsl.Attribute("completed", dsl.Boolean, "Is the todo completed", func() {})
	dsl.Required("id", "title", "completed")
})

var _ = dsl.Service("todo", func() {
	dsl.Description("The todo service performs operations on Todos.")
	dsl.Method("hello", func() {
		dsl.Payload(func() {
			dsl.Attribute("name", dsl.String, "Name")
			dsl.Required("name")
		})
		dsl.Result(dsl.String)
		dsl.HTTP(func() {
			dsl.GET("/hello/{name}")
			dsl.Response(dsl.StatusOK)
		})
	})

	dsl.Method("create", func() {
		dsl.Payload(func() {
			dsl.Attribute("title", dsl.String)
			dsl.Required("title")
		})
		dsl.Result(Todo)
		dsl.HTTP(func() {
			dsl.POST("/todos")
			dsl.Response(dsl.StatusCreated)
		})
	})
})
