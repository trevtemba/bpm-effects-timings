# **BPM Effects Timings CLI**  

A simple CLI tool that calculates baseline effect timings (Reverb Delay, Pre-Delay, Compressor Attack/Release) based on BPM, I got sick of using a calculator to find those pocket-timings!

~*Trevor Katemba*. 

## **Installation**  

### **Windows**  
1. **Download the latest release**  
   - Go to the [Releases](https://github.com/yourusername/bpm-effects-timings/releases) page.  
   - Download `t.exe`.  

2. **Move the executable to a directory in your PATH**  
   - Open **PowerShell** and run:  
     ```powershell
     move t.exe C:\Windows\System32\
     ```
   - Now, you can use `t` from any command prompt or terminal.  

---

### **Linux & macOS**  
1. **Download the latest release**  
   ```sh
   curl -L -o t https://github.com/yourusername/bpm-effects-timings/releases/latest/download/t-linux
   ```
   (For macOS, replace t-linux with t-mac in the URL above.)

2. **Make it executable and move it to /usr/local/bin/**
   ```sh
   chmod +x t
   sudo mv t /usr/local/bin/

## **Usage**
Once you've installed it, run the command in your CLI with a BPM value:
```sh
t 120
```
If no BPM is provided, the tool will just prompt you to enter one manually.

**Example Output for BPM of 100**
```sh
Effects Timings:
----------------------------
Reverb Delay:         2.40s
Reverb Pre-Delay:    37.50ms
Compressor Attack:   37.50ms
Compressor Release:  75.00ms
----------------------------
```

### **Building with Go** ###
If you have Go installed, you can also build the tool manually :)
```sh
git clone https://github.com/yourusername/bpm-effects-timings.git
cd bpm-effects-timings
go build -o t ./cmd/main.go
sudo mv t /usr/local/bin/   # Linux/macOS
move t.exe C:\Windows\System32\   # Windows
```

### **License**
This project is under MIT license, feel free to modify and contribute to it!

*Thanks for using!*
