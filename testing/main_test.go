package testing

import (
	"testing"

	"github.com/preslavmihaylov/todocheck/testing/config"
)

func TestSum(t *testing.T) {
	err := config.NewScenario().
		WithBinary("../todocheck").
		WithBasepath("./scenarios/simple_todos").
		WithConfig("./scenarios/simple_todos/.todocheck.yaml").
		Run()
	if err != nil {
		t.Errorf("todocheck scenario failed: %w", err)
	}
}
