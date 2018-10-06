package twofer

import "fmt"

// ShareWith replaces text with "you" if you don't pass anything or with the name you pass
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
