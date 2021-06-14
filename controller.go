package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

type Controller struct {
	Storage     Storage
	Template    *template.Template
	NameRegexp  *regexp.Regexp
	EmailRegexp *regexp.Regexp
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	members := c.Storage.GetAllMembers()
	c.Template.Execute(w, members)
}

func (c *Controller) Save(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	if !c.NameIsValid(name) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: name is not valid\n")
		return
	}
	if !c.EmailIsValid(email) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: email is not valid\n")
		return
	}
	if err := c.Storage.SaveMember(name, email); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v\n", err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func CreateController() *Controller {
	return &Controller{
		Storage:     CreateSliceStorage(),
		Template:    template.Must(template.ParseFiles("index.html")),
		NameRegexp:  regexp.MustCompile(`^([A-z]+\.?\s?)+$`),
		EmailRegexp: regexp.MustCompile(`^[A-z\.0-9]+@[A-z\.0-9]+$`),
	}
}

func (c *Controller) NameIsValid(name string) bool {
	return c.NameRegexp.MatchString(name)
}

func (c *Controller) EmailIsValid(email string) bool {
	return c.EmailRegexp.MatchString(email)
}
