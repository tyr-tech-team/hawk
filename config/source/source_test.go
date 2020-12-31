package source

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tyr-tech-team/hawk/pkg/consul"
)

func createFile() {
	data := []byte("{\"name\": \"jame\" , \"age\":10}")
	f, err := os.Create("test.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}

	f.Sync()

}

func deleteFile() {
	err := os.Remove("test.json")
	if err != nil {
		panic(err)
	}

}

func Test_NewFileSource(t *testing.T) {
	createFile()
	defer deleteFile()

	s := NewFile("test.json")
	data, err := s.Read()
	assert.NoError(t, err)

	t.Log(string(data.Data))
}

func Test_NewConsulSource(t *testing.T) {
	cli := consul.NewClient(context.TODO(), consul.Config{
		Address: "localhost:8500",
	})

	err := cli.Set("test", []byte("name: \"john\"\nage: 123\n"))
	assert.NoError(t, err)

	s := NewConsul(cli, "test")
	d, err := s.Read()
	assert.NoError(t, err)
	assert.NotEmpty(t, d)
}
