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
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/psidh/Ganges/src/repl"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error: Unable to retrieve user information:", err)
		panic(err)
	}

	fmt.Println("🌊 Welcome to the Ganges Programming Language, " + currentUser.Username + "!")
	fmt.Println("The Ganges language is designed to empower developers with simplicity and power.")
	fmt.Println("Feel free to write your code using our expressive syntax, and let the code flow like a river!")

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("Current session started at:", currentTime)

	fmt.Println("\n✍️ Type your code below:")

	repl.Start(os.Stdin, os.Stdout)

	fmt.Println("\n🌊 Ganges has completed your code execution. Have a productive day!")
}
