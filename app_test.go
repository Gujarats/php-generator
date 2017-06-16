package main

import (
	"fmt"
	"testing"
)

func TestGetSetterGetter(t *testing.T) {
	testObjects := []struct {
		Input    string
		Prefix   string
		Expected string
	}{
		{
			Input:  "myVar",
			Prefix: "public",
		},

		{
			Input:  "room",
			Prefix: "private",
		},
	}

	for _, testObject := range testObjects {
		result := getSetterGetter(testObject.Prefix, testObject.Input)
		fmt.Println(result)
	}
}

func TestGetVariables(t *testing.T) {
	testObjects := []struct {
		Datas    []byte
		Expected []string
	}{
		{
			Datas: []byte(`<?php
				class SomeClass {
    				private $myVar;
    				private $newVaria;
    				private $newVar;

				}`),
			Expected: []string{"myVar", "newVaria", "newVar"},
		},

		{
			Datas: []byte(`<?php
				class SomeClass {
    				private $hello;
    				private $Wkwkw;
    				private $eheHaha;

				}`),
			Expected: []string{"hello", "Wkwkw", "eheHaha"},
		},
	}

	for _, testObject := range testObjects {
		variables, indexEndClass := getVarNames(testObject.Datas)
		if indexEndClass <= 0 {
			t.Errorf("End class not found")
		}
		if !eq(variables, testObject.Expected) {
			t.Errorf("actual = %v, expected = %v\n", variables, testObject.Expected)
		}

	}
}

func TestGetArguments(t *testing.T) {
	testObjects := []struct {
		Input    []string
		Expected string
	}{
		{
			Input:    []string{"a", "b", "c"},
			Expected: "$a,$b,$c",
		},
		{
			Input:    []string{"a", "b", "haha"},
			Expected: "$a,$b,$haha",
		},
	}

	for _, testObject := range testObjects {
		actual := getArgments(testObject.Input)
		if actual != testObject.Expected {
			t.Errorf("actual = %v, expected = %v\n", actual, testObject.Expected)
		}
	}
}

func TestGetConstructor(t *testing.T) {
	testObjects := []struct {
		Argguments []string
		Expected   string
	}{
		{
			Argguments: []string{"a", "b", "c"},
		},

		{
			Argguments: []string{"id", "name", "price"},
		},
	}

	for _, testObject := range testObjects {
		actual := constructor(testObject.Argguments)
		fmt.Println(actual)
	}
}

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if a == nil || b == nil {
		return false
	}

	for index, value := range a {
		if value != b[index] {
			return false
		}
	}

	return true

}
