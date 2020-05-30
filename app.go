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
	"fmt"
	"os"
	"net/http"
  "log"
  "io/ioutil"
)

// Logging
const (
		// Logs on/off
    VERBOSE bool = false
)

// Messages to returned
const (
		// Message for status code 200
    REGULAR_MESSAGE 								string = "Regular CPF (OK)"

		// Message for status code 400
    INVALID_CPF_FORMAT_MESSAGE 			string = "Invalid CPF format"

		// Message for status code 403
		SUBJECT_REJECTED_MESSAGE 				string = "CPF not regular or not existant"

		// Message for status code 500
		EXTERNAL_SERVER_ERROR_MESSAGE 	string = "Server communication problem"

		// Message for any other status code
		UNKNOWN_ERROR_MESSAGE 					string = "Unknown error"

		// Message for bad arguments
		INVALID_PARAMETER_MESSAGE 			string = "Invalid arguments, call with 1 and only 1 CPF"
)

// Main function
func main() {

	if VERBOSE == false {
		log.SetOutput(ioutil.Discard) // Turning logs OFF
	}

	args := os.Args[1:]

	if params := len(args); params != 1 {
		fmt.Println("CPF:      ?")
		fmt.Println("SUCCESS:  False")
		fmt.Println("MESSAGE: ", INVALID_PARAMETER_MESSAGE)
		os.Exit(1)
	}
	fmt.Println("CPF:     ", args[0])

	cpfValidationCode := validateCpf(args[0])

	// Creating Response...

	var status string
	var message string

	switch cpfValidationCode {
		case http.StatusOK:
			status = "True"
			message = REGULAR_MESSAGE
		case http.StatusForbidden:
			status = "False"
			message = SUBJECT_REJECTED_MESSAGE
		case http.StatusBadRequest:
			status = "False"
			message = INVALID_CPF_FORMAT_MESSAGE
		case http.StatusInternalServerError:
			status = "False"
			message = EXTERNAL_SERVER_ERROR_MESSAGE
		default:
			status = "False"
			message = UNKNOWN_ERROR_MESSAGE
	}

	fmt.Println("SUCCESS: ", status)
	fmt.Println("MESSAGE: ", message)
	os.Exit(0)
}
