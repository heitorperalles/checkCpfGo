//------------------------------------------------------------------------------
// From http://github.com/heitorperalles/checkCpfGo
//
// Distributed under The MIT License (MIT) <http://opensource.org/licenses/MIT>
//
// Copyright (c) 2020 Heitor Peralles <heitorgp@gmail.com>
//
// Permission is hereby  granted, free of charge, to any  person obtaining a copy
// of this software and associated  documentation files (the "Software"), to deal
// in the Software  without restriction, including without  limitation the rights
// to  use, copy,  modify, merge,  publish, distribute,  sublicense, and/or  sell
// copies  of  the Software,  and  to  permit persons  to  whom  the Software  is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE  IS PROVIDED "AS  IS", WITHOUT WARRANTY  OF ANY KIND,  EXPRESS OR
// IMPLIED,  INCLUDING BUT  NOT  LIMITED TO  THE  WARRANTIES OF  MERCHANTABILITY,
// FITNESS FOR  A PARTICULAR PURPOSE AND  NONINFRINGEMENT. IN NO EVENT  SHALL THE
// AUTHORS  OR COPYRIGHT  HOLDERS  BE  LIABLE FOR  ANY  CLAIM,  DAMAGES OR  OTHER
// LIABILITY, WHETHER IN AN ACTION OF  CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE  OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//------------------------------------------------------------------------------
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
