package healthz

import (
	"testing"
)

/*
Health check Usecase UnitTest
*/
func TestExecute(t *testing.T) {
	providers := []struct {
		name     string
		expected bool
	}{
		{
			name:     "正常系",
			expected: true,
		},
	}

	for _, provider := range providers {
		actual := Execute()
		if provider.expected != actual {
			t.Errorf("provider %s: expected %t, actual %t", provider.name, provider.expected, actual)
		}
	}
}
