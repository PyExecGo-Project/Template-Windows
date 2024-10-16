# 🐍 PyExecGo Project Template for Windows

> **Note**: This is a template where most of the setup is already done to use PyExecGo. I’m working on an automatic builder that will handle these steps automatically! :)

PyExecGo is a simple tool that allows you to run Python programs without needing Python installed on the target machine. It uses a portable version of Python and automatically installs any dependencies listed in `requirements.txt`.

## 🛠️ Getting Started

Follow the steps below to set up and run your Python program using PyExecGo:

### Prerequisites

- Go installed on your system (https://go.dev/dl/)
- A Python program that you want to run

### Installation Steps

1. **Clone the Repository**

   Clone this repository to your local machine:

   ```bash
   git clone https://github.com/PyExecGo-Project/template.git
   cd template
   ```

2. **Organize Your Files**

   Place all your Python files in the same folder as `main.go`.

3. **Insert the Special Sauce**

   Before compiling, **copy the contents of `special-sauce.py` and paste it at the beginning of your `main.py` (or your start script).**  
   **This code must be on line 1 of your Python script.**

   After copying the content, you can safely delete `special-sauce.py`.

4. **Edit `main.go`**

   Open `main.go` and make any necessary modifications to fit your project’s needs.

5. **Build the Executable**

   Build the Go executable:

   ```bash
   go build main.go
   ```

6. **Cleanup (Optional)**

   If you wish, you can delete `main.go` and `go.mod` after building the executable to keep things tidy.

7. **Run Your Program**

   You can now run your Python program using the generated `main.exe`.

   When you run `main.exe`, it will automatically install the dependencies listed in `requirements.txt`.

## 📝 Notes

- Ensure that `requirements.txt` is present in the same folder as your Python files.
- You can customize the startup process by modifying the Go code in `main.go` before building.

## 🤝 Contributing

Feel free to contribute by submitting a pull request or opening an issue.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/PyExecGo-Project/template/blob/main/LICENSE) file for details.
