package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Arfan")
	}
}

func TestTableHelloWorld(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name : "Arfan",
			request : "Arfan",
			expected : "Hello Arfan",
		},
		{
			name : "Hidayatullah",
			request : "Hidayatullah",
			expected : "Hello Hidayatullah",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T){
		result := HelloWorld("Arfan")

		require.Equal(t, "Hello Arfan", result, "Result must be 'Hello Alex'")
	})
	t.Run("Kurniawan", func(t *testing.T){
		result := HelloWorld("Hidayatullah")

		require.Equal(t, "Hello Hidayatullah", result, "Result must be 'Hello Hidayatullah'")
	})
}

func TestMain(m *testing.M){
	//before 
	fmt.Println("Before unit test")

	m.Run()

	//after
	fmt.Println("After unit test")
}

func TestHelloWorldEko(t *testing.T){
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		//error
		t.Fail()
	}

	fmt.Println("ini TestHelloWorldEko Done")
}

func TestHelloWorldKhannedy(t *testing.T){
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		//error
		t.FailNow()
	}
	fmt.Println("ini TestHelloWorldKhannedy")
}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Arfan")

	require.Equal(t, "Hello Arfan", result, "Result must be 'Hello Alex'")
	
	fmt.Println("ini TestHelloWorldArfan")
}

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Alex")

	assert.Equal(t, "Hello Alex", result, "Result must be 'Hello Alex'")

	fmt.Println("ini TestHelloWorldAlex with Assert done")
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "darwin" {
		t.Skip("can not run on Mac OS")
	}

	result := HelloWorld("Arfan")
	require.Equal(t, "Hello Arfan", result, "Result must be 'Hello Arfan'")
}