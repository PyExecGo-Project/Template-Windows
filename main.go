package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println(`  

	_____       ______                _____       
	|  __ \     |  ____|              / ____|      
	| |__) |   _| |__  __  _____  ___| |  __  ___  
	|  ___/ | | |  __| \ \/ / _ \/ __| | |_ |/ _ \ 
	| |   | |_| | |____ >  <  __/ (__| |__| | (_) |
	|_|    \__, |______/_/\_\___|\___|\_____|\___/ 
			__/ |                                  
		   |___/                                   
   
   
	`)
	fmt.Println("© PyExecGo Contributors - PyExecGo Executable version: v1.1.0")
	fmt.Println("PyExecGo is released under the MIT License")
	fmt.Println("Copyright and license of original developers are respected and can be found in LICENSE.md if applicable.")
	fmt.Println("This executable was built for the project: github.com/PyExecGo-Project/template")
	fmt.Println()

	// Show splash screen for 5 seconds
	for i := 5; i > 0; i-- {
		fmt.Printf("\rStarting in %d seconds...", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()

	// Change the current working directory to portable-python-bin
	if err := os.Chdir("portable-python-bin"); err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Set environment variables
	if err := os.Setenv("PYTHONHOME", ".\\"); err != nil {
		fmt.Println("Error setting PYTHONHOME:", err)
		return
	}
	if err := os.Setenv("PYTHONUNBUFFERED", "1"); err != nil {
		fmt.Println("Error setting PYTHONUNBUFFERED:", err)
		return
	}

	// Check if requirements.txt exists before installing dependencies
	if _, err := os.Stat("../requirements.txt"); err == nil {
		// Installing dependencies
		installCmd := exec.Command(".\\python.exe", "-m", "pip", "install", "-r", "..\\requirements.txt")
		if err := installCmd.Run(); err != nil {
			fmt.Println("Error installing dependencies:", err)
			return
		}
	} else {
		fmt.Println("requirements.txt not found. Skipping installation.")
	}

	// Execute the main.py script
	fmt.Println("Executing:", "main.py")
	runCmd := exec.Command(".\\python.exe", "..\\main.py")

	// Create a pipe to read stdout
	stdout, err := runCmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe:", err)
		return
	}

	if err := runCmd.Start(); err != nil {
		fmt.Println("Error starting main.py:", err)
		return
	}

	// Print output in real-time
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdout:", err)
	}

	if err := runCmd.Wait(); err != nil {
		fmt.Println("Error waiting for main.py to finish:", err)
	}

	fmt.Println("main.py executed successfully.")
}
