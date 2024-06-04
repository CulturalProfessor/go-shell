# My Shell

A simple command-line shell implemented in Go. This shell supports a variety of built-in commands and can execute external programs. It includes basic features like navigating directories, printing the current working directory, and handling commands with arguments.

## Features

- **Built-in Commands**:
  - `exit 0`: Exits the shell.
  - `echo <text>`: Prints the provided text to the standard output.
  - `pwd`: Prints the current working directory.
  - `cd <directory>`: Changes the current working directory to the specified directory. Supports absolute paths, relative paths, and the `~` character for the home directory.
  - `type <command>`: Indicates whether a command is a shell built-in or an external command.

- **External Commands**:
  - Supports execution of any external command available in the directories specified in the `PATH` environment variable.

## Usage

### Building the Shell

1. Clone the repository:
    ```sh
    git clone https://github.com/your-username/my-shell.git
    cd my-shell
    ```

2. Build the executable:
    ```sh
    go build -o my_shell
    ```

### Running the Shell

Run the built executable:
```sh
./my_shell
