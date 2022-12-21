package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			result: "Hello, Paulo!",
			items:  []string{"Paulo"},
		},
		{
			result: "Hello, Lucas, Celeste, Isaldo!",
			items:  []string{"Lucas", "Celeste", "Isaldo"},
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("expected %s, got %s", s, st.result)
		}
	}
}
