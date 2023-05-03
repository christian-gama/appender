# File Appender
File Appender is a command-line utility written in Go that reads files with specified extensions in a given directory and appends their content into a single file called "output.txt". The program can also exclude specified directories while searching for files.
This is useful to merge files and feed them into an AI model.

# Installation
First, make sure you have Go installed on your system.
Clone the repository or download the main.go file.
Navigate to the directory containing main.go and run the following command to build the executable:
```make build```

# Usage
Run the program with the required arguments:
```./appender <file_extensions> <directory> [excluded_directories]```

- <file_extensions>: A comma-separated list of file extensions (e.g., ".txt,.md,.py").
- <directory>: The path to the directory where the program should search for files (e.g., "/path/to/your/directory").
- [excluded_directories]: (Optional) A comma-separated list of directory names to exclude while searching for files (e.g., "excluded1,excluded2").

# Output
The program creates a file named "output.txt" in the same directory as the executable. The content of each file found in the specified directory is appended to "output.txt", along with the file path in the format "# /path/to/the/file.py" at the beginning of the appended content.

