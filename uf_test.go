package brasil

import (
	"testing"
)

func TestUF(t *testing.T) {
	var result string
	var err error

	result, err = Estado("SP")
	if err != nil {
		t.Fatal(err)
	}
	if result != "São Paulo" {
		t.Fatal(result)
	}

	result, err = Estado("MS")
	if err != nil {
		t.Fatal(err)
	}
	if result != "Mato Grosso do Sul" {
		t.Fatal(result)
	}
}

func TestEstado(t *testing.T) {
	var result string
	var err error

	result, err = UF("São PAULO")
	if err != nil {
		t.Fatal(err)
	}
	if result != "SP" {
		t.Fatal(result)
	}

	result, err = UF("AmApá")
	if err != nil {
		t.Fatal(err)
	}
	if result != "AP" {
		t.Fatal(result)
	}
}
