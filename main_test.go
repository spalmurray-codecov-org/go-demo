package main

import "testing"

func TestAdd(t *testing.T) {
    if add(1, 2) != 3 {
        t.Fatalf("add(1, 2) != 3")
    }
}

func TestSubtract(t *testing.T) {
    if subtract(2, 1) != 1 {
        t.Fatalf("subtract(2, 1) != 1")
    }
}

func TestMultiply(t *testing.T) {
    if multiply(2, 3) != 6 {
        t.Fatalf("multiply(2, 3) != 6")
    }
}
