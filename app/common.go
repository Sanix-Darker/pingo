package main

import "fmt"

func Format(anyThing ...interface{}) string {
	return fmt.Sprintf(
		anyThing[0].(string),
		anyThing[1:]...,
	)
}
