package main

import (
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
	fmt.Println("© PyExecGo Contributors - PyExecGo Executable version: v1.2.0")
	fmt.Println("PyExecGo is released under the MIT License")
	fmt.Println("Copyright and license of original developers are respected and can be found in LICENSE.md if applicable.")
	fmt.Println("This executable was built for the project: github.com/PyExecGo-Project/template")
	fmt.Println()

	for i := 1; i > 0; i-- {
		fmt.Printf("\rStarting in %d seconds...", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()

	if err := os.Chdir("portable-python-bin"); err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	if err := os.Setenv("PYTHONHOME", ".\\"); err != nil {
		fmt.Println("Error setting PYTHONHOME:", err)
		return
	}
	if err := os.Setenv("PYTHONUNBUFFERED", "1"); err != nil {
		fmt.Println("Error setting PYTHONUNBUFFERED:", err)
		return
	}

	if _, err := os.Stat("../requirements.txt"); err == nil {
		installCmd := exec.Command(".\\python.exe", "-m", "pip", "install", "-r", "..\\requirements.txt")
		if err := installCmd.Run(); err != nil {
			fmt.Println("Error installing dependencies:", err)
			return
		}
	} else {
		fmt.Println("requirements.txt not found. Skipping installation.")
	}

	fmt.Println("Opening a new command prompt to run the Python script...")

	cmd := exec.Command("cmd", "/C", "start", "cmd", "/K", ".\\python.exe ..\\main.py")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error opening new command prompt:", err)
		return
	}

	fmt.Println("Python script is now running in a new command prompt. Exiting Go program.")
}
