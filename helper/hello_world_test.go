package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "ifqy",
			request: "ifqy",
		},

		{
			name:    "gifha",
			request: "gifha",
		},
		{
			name:    "azhar",
			request: "azhar",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}

func BenchmarkSub(b *testing.B) {
	b.Run("azhar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("azhar")
		}
	})

	b.Run("ifqy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("ifqy")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("ifqy")
	}
}
func BenchmarkHelloWorldAzhar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Azhar")
	}
}

func TestTable(t *testing.T) {
	test := []struct {
		Name      string
		Request   string
		Ekspetasi string
	}{
		{

			Name:      "azhar",
			Request:   "azhar",
			Ekspetasi: "hello azhar",
		},

		{
			Name:      "gifha",
			Request:   "gifha",
			Ekspetasi: "hello gifha",
		},
	}

	for _, test := range test {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			require.Equal(t, test.Ekspetasi, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("azhar", func(t *testing.T) {
		result := HelloWorld("azhar")
		require.Equal(t, "hello azhar", result, "result must be 'hello'")
	})

	t.Run("gifha", func(t *testing.T) {
		result := HelloWorld("gifha")
		require.Equal(t, "hello gifha", result, "result must be 'hello'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("before")

	m.Run()

	fmt.Println("after")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("os nya linux")
	}
	result := HelloWorld("azhar")
	require.Equal(t, "Hello azhar", result, "result must be 'hello'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("azhar")
	require.Equal(t, "Hello azhar", result, "result must be 'hello'")

	fmt.Println("eksekusi")
}

func TestHelloWorldAsset(t *testing.T) {
	result := HelloWorld("azhar")
	assert.Equal(t, "Hello azhar", result, "result must be 'hello'")

	fmt.Println("eksekusi")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("azhar")

	if result != "Hello azhar" {
		//error
		t.Fatal("harusnya 'hello'")
	}
}
