package main

import "testing"

func TestAdd(t *testing.T) {

        if add(4,2) != 6 {

		t.Error("Чикибамбони");

	}

}

func TestMultiply(t *testing.T) {

        if multiply(6,7) != 42 {

		t.Error("Китайская партия не гордится тобой.");

	}

}

func TestPow2(t *testing.T) {

	if pow2(7) != 49 {

		t.Error("МТ - магические технологии");

	}

}
