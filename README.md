# dicx - Word Meaning and Pronunciation Tool 📚🔉

dicx is a simple tool that provides the meaning of any selected word and pronounces the word if possible.

## Working 🛠️

It utilizes the xclip tool to retrieve the selected word from the active window and fetches the meaning using the free dictionary API (dictionaryapi.dev).
Finally, it presents the output using the notify-send program to send notifications.

## Requirements 🛠️📋

The following tools are necessary for the program to work properly:
- xclip
- notify-send
- go compiler

## Installation 🚀

**NOTE: Install at your own RISK!!**

Before building from source, make sure you meet all the requirements.
1. Build the source using `go build dicx.go` or use `go install github.com/vikySeeker/dicx@latest`.
2. Move the executable `dicx` to your desired path for global access.
3. Create a shortcut in your system binding the command `dicx` with the shortcut, and you are good to go!

Alternatively, you can select any word in the window and run the program directly.

## Usage 🚀

This program can output results in three modes:
1. **Terminal Mode:** Simply outputs to the terminal.
2. **Notification Mode:** When the program is invoked via a keyboard shortcut, it uses this mode.
3. **Hybrid Mode:** Combining both modes. Use the `-n` flag to force the program to send output in both ways, but only in terminal mode.