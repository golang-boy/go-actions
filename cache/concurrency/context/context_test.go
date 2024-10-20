package context

import (
	"context"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	ctx := context.Background()

	parent := context.WithValue(ctx, "key", "parent value")

	child := context.WithValue(parent, "key", "child value")

	t.Log("parent value: ", parent.Value("key"))
	t.Log("child value: ", child.Value("key"))

	child2, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()

	// child2中携带parent的值
	t.Log("child2 value: ", child2.Value("key"))

	child3 := context.WithValue(parent, "new key", "child3 value")

	t.Log("parent value: ", parent.Value("new key"))
	t.Log("child3 value: ", child3.Value("new key"))

	parent1 := context.WithValue(parent, "map", map[string]string{})
	child4, cancel := context.WithTimeout(parent1, 1*time.Second)
	defer cancel()

	m := child4.Value("map").(map[string]string)
	m["key"] = "child4 value"

	nm := parent1.Value("map").(map[string]string)
	t.Log("parent value: ", nm["key"])

}
