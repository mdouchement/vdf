package vdf_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/mdouchement/vdf"
)

var remotecache = `"429660"
{
	"Save000.tobsav"
	{
		"root"		"0"
		"size"		"718088"
		"localtime"		"1520277506"
		"time"		"1520277515"
		"remotetime"		"1520277515"
		"sha"		"e0c260324b8a3119629325ab5b8c4420da5e7ade"
		"syncstate"		"1"
		"persiststate"		"0"
		"platformstosync2"		"-1"
	}
	"quicksave.tobsav"
	{
		"root"		"0"
		"size"		"718088"
		"localtime"		"1524340445"
		"time"		"1524340447"
		"remotetime"		"1524340447"
		"sha"		"1b2c5f42828c1f578b05f48ac75a2d031336cc3d"
		"syncstate"		"1"
		"persiststate"		"0"
		"platformstosync2"		"-1"
	}
}
`

func TestParse(t *testing.T) {
	remote, err := vdf.Parse([]byte(remotecache))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	iface, ok := remote["429660"]
	if !ok {
		t.Fatalf("Expected 429660 to be present")
	}

	saves, ok := iface.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected saves to be casted in map[string]interface{}")
	}

	_, ok = saves["Save000.tobsav"]
	if !ok {
		t.Fatalf("Expected Save000.tobsav to be present")
	}

	iface, ok = saves["quicksave.tobsav"]
	if !ok {
		t.Fatalf("Expected quicksave.tobsav to be present")
	}

	quicksave, ok := iface.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected quicksave to be casted in map[string]interface{}")
	}

	if root, ok := quicksave["root"]; !ok || root.(string) != "0" {
		t.Fatalf("Expected quicksave[\"root\"] to be present and have value 0")
	}

	if localtime, ok := quicksave["localtime"]; !ok || localtime.(string) != "1524340445" {
		t.Fatalf("Expected quicksave[\"localtime\"] to be present and have value 1524340445")
	}

	if sha, ok := quicksave["sha"]; !ok || sha.(string) != "1b2c5f42828c1f578b05f48ac75a2d031336cc3d" {
		t.Fatalf("Expected quicksave[\"sha\"] to be present and have value 1b2c5f42828c1f578b05f48ac75a2d031336cc3d")
	}
}

func TestParseIO(t *testing.T) {
	r := bytes.NewBufferString(remotecache)
	remote, err := vdf.ParseIO(r)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	_, ok := remote["429660"]
	if !ok {
		t.Fatalf("Expected 429660 to be present")
	}
}

func TestGenerate(t *testing.T) {
	remote, err := vdf.Parse([]byte(remotecache))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	tmp, err := vdf.Generate(remote)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	remote2, err := vdf.Parse(tmp)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !reflect.DeepEqual(remote, remote2) {
		t.Fatalf("Expected\n%v\nto equal\n%v", remote2, remote)
	}
}

func TestGenerateIO(t *testing.T) {
	remote, err := vdf.Parse([]byte(remotecache))
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	w := new(bytes.Buffer)
	if err = vdf.GenerateIO(w, remote); err != nil {
		t.Fatalf("Error: %v", err)
	}

	remote2, err := vdf.Parse(w.Bytes())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if !reflect.DeepEqual(remote, remote2) {
		t.Fatalf("Expected\n%v\nto equal\n%v", remote2, remote)
	}
}
