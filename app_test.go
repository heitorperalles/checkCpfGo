package main

import (
		"log"
		"testing"
		"net/http"
)

// Main TEST Function
//
// This function is executed before all tests
//
// Param according to testing package pattern [func TestMain(m *testing.M)]
func TestMain(m *testing.M) {

		log.Println("Starting TESTS...")
		m.Run()
}

// Test validateCpf function
//
// This test consists of calling the function by passing a CPF without numbers
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfLetters(t *testing.T) {

	log.Println("Testing validateCpf passing a CPF without numbers...")

	err := validateCpf("INVALID")

	if err != http.StatusBadRequest {
		t.Errorf("The function returned a code different than BadRequest: %d", err)
	} else {
		log.Println("Success by calling validateCpf with a CPF without numbers.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing an empty CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfEmpty(t *testing.T) {

	log.Println("Testing validateCpf passing an empty CPF...")

	err := validateCpf("")

	if err != http.StatusBadRequest {
		t.Errorf("The function returned a code different than BadRequest: %d", err)
	} else {
		log.Println("Success by calling validateCpf with an empty CPF.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing a malformed  CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfMalformed(t *testing.T) {

	log.Println("Testing validateCpf passing a malformed CPF...")

	err := validateCpf("1234567")

	if err != http.StatusBadRequest {

		// Not rising an error because this test depends on SERPRO API response

		t.Skipf("The function returned a code different than BadRequest: %d", err)
	} else {
		log.Println("Success by calling validateCpf with a malformed CPF.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing a canceled  CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfCanceled(t *testing.T) {

	log.Println("Testing validateCpf passing a canceled CPF...")

	err := validateCpf("64913872591")

	if err != http.StatusForbidden {

		// Not rising an error because this test depends on SERPRO API response

		t.Skipf("The function returned a code different than Forbidden: %d", err)
	} else {
		log.Println("Success by calling validateCpf with a canceled CPF.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing an inexistant CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfInexistant(t *testing.T) {

	log.Println("Testing validateCpf passing an inexistant CPF...")

	err := validateCpf("11334739706")

	if err != http.StatusForbidden {

		// Not rising an error because this test depends on SERPRO API response

		t.Skipf("The function returned a code different than Forbidden: %d", err)
	} else {
		log.Println("Success by calling validateCpf with an inexistant CPF.")
	}
}

// Test validateCpf function
//
// This test consists of calling the function by passing a regular CPF
//
// Param according to testing package [func TestXxx(*testing.T)]
func TestValidateCpfRegular(t *testing.T) {

	log.Println("Testing validateCpf passing a regular CPF...")

	err := validateCpf("40442820135")

	if err != http.StatusOK {

		// Not rising an error because this test depends on SERPRO API response

		t.Skipf("The function returned a code different than OK: %d", err)
	} else {
		log.Println("Success by calling validateCpf with a regular CPF.")
	}
}
