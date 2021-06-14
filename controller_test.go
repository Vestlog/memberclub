package main

import (
	"testing"
)

func TestNameIsValid(t *testing.T) {
	c := CreateController()
	data := []string{
		"John Snow",
		"Dana A. Watkins",
		"Mr. Brown",
	}
	for _, entry := range data {
		if !c.NameIsValid(entry) {
			t.Errorf("entry should be valid: %s", entry)
		}
	}
}

func TestNameIsValidFalse(t *testing.T) {
	c := CreateController()
	data := []string{
		"1337pwn",
		"mail@example.com",
		"",
	}
	for _, entry := range data {
		if c.NameIsValid(entry) {
			t.Errorf("entry should not be valid: %s", entry)
		}
	}
}

func TestEmailIsValid(t *testing.T) {
	c := CreateController()
	data := []string{
		"mail@example.com",
		"first.entry2000@example.com",
	}
	for _, entry := range data {
		if !c.EmailIsValid(entry) {
			t.Errorf("email should be valid: %s", entry)
		}
	}
}

func TestEmailIsValidFalse(t *testing.T) {
	c := CreateController()
	data := []string{
		"example.com",
		"@zxczxc@",
	}
	for _, entry := range data {
		if c.EmailIsValid(entry) {
			t.Errorf("email should not be valid: %s", entry)
		}
	}
}
