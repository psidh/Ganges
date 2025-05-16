// Copyright 2025 PHILKHANA SIDHARTH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/psidh/Ganges/src/eval"
	"github.com/psidh/Ganges/src/lexer"
	"github.com/psidh/Ganges/src/object"
	"github.com/psidh/Ganges/src/parser"
	"github.com/rs/cors"
)

type RequestPayload struct {
	Code string `json:"code"`
}

type ResponsePayload struct {
	Output string `json:"output"`
}

func executeGangesCode(code string) string {
	var outputBuffer bytes.Buffer

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		w.Close()
		os.Stdout = oldStdout

		outputBuffer.WriteString("🛑 Parser errors:\n")
		for _, msg := range p.Errors() {
			outputBuffer.WriteString("  - " + msg + "\n")
		}
		return outputBuffer.String()
	}

	env := object.NewEnvironment()
	evaluated := eval.Eval(program, env)

	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	os.Stdout = oldStdout

	capturedOutput := buf.String()
	if capturedOutput != "" {
		outputBuffer.WriteString("Console output:\n" + capturedOutput + "\n\n")
	}

	if evaluated != nil {
		outputBuffer.WriteString("Result: \n" + evaluated.Inspect())
		return outputBuffer.String()
	}

	if capturedOutput == "" {
		return "✅ Program executed successfully (no output)"
	}

	return outputBuffer.String()
}

func codeExecutionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqPayload RequestPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqPayload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	output := executeGangesCode(reqPayload.Code)

	respPayload := ResponsePayload{
		Output: output,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respPayload)
}

func main() {
	logger := log.New(os.Stdout, "[Ganges Server] ", log.LstdFlags)

	mux := http.NewServeMux()
	mux.HandleFunc("/execute", codeExecutionHandler)

	handler := cors.Default().Handler(mux) // Enable CORS

	port := os.Getenv("PORT")
	host := "localhost"

	if port == "" {
		port = "3001"
	} else {
		host = "0.0.0.0" // For Render
	}

	addr := host + ":" + port
	logger.Printf("Server started at http://%s", addr)
	logger.Fatal(http.ListenAndServe(addr, handler))
}
