package main

import (
	"testing"
)

func TestEvaluatePrefix_SimpleExpressions(t *testing.T) {
	tests := []struct {
		expression string
		expected   int
	}{
		{"+ 2 3", 5},
		{"- 10 4", 6},
		{"* 3 7", 21},
		{"/ 8 2", 4},
		{"^ 2 3", 8},
	}

	for _, tt := range tests {
		result, err := EvaluatePrefix(tt.expression)
		if err != nil {
			t.Errorf("unexpected error for %q: %v", tt.expression, err)
		}
		if result != tt.expected {
			t.Errorf("EvaluatePrefix(%q) = %d; want %d", tt.expression, result, tt.expected)
		}
	}
}

func TestEvaluatePrefix_ComplexExpression(t *testing.T) {
	expr := "+ 5 * - 4 2 ^ 3 2" 
	expected := 23

	result, err := EvaluatePrefix(expr)
	if err != nil {
		t.Errorf("unexpected error for complex expression: %v", err)
	}
	if result != expected {
		t.Errorf("EvaluatePrefix(%q) = %d; want %d", expr, result, expected)
	}
}

func TestEvaluatePrefix_InvalidExpressions(t *testing.T) {
	invalidInputs := []string{
		"",            
		"+ 5",         
		"+ 5 a",       
		"/ 5 0",       
		"^ 2 -3",      
		"+ 5 *",       
	}

	for _, expr := range invalidInputs {
		_, err := EvaluatePrefix(expr)
		if err == nil {
			t.Errorf("expected error for invalid input %q, but got none", expr)
		}
	}
}


func ExampleEvaluatePrefix() {
	result, err := EvaluatePrefix("+ 5 * - 4 2 ^ 3 2")
	if err != nil {
		panic(err)
	}
	println(result)
	
}
