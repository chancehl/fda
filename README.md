# **Fast Directory Alias (fda)**

`fda` is a simple CLI utility that allows users to create shell aliases for directories, making navigation easier. Running `fda <name>` adds an alias to your shell's configuration file, allowing you to quickly `cd` into frequently used directories.

## **Features**

- ✅ Creates a shell alias for the current directory with a user-friendly name.
- ✅ Supports `bash` and `zsh` by updating `.bashrc` or `.zshrc`.
- ✅ Automatically detects the user's shell.
- ✅ Validates alias names to prevent invalid or duplicate entries.
- ✅ Provides a fallback to the user's home directory if no directory is specified.

## **Installation**

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/fda.git
   cd fda
   ```
2. Build the binary:
   ```sh
   go build -o fda
   ```
3. Move it to a directory in your `PATH` (optional):
   ```sh
   mv fda /usr/local/bin/
   ```

## **Usage**

### **Adding an alias for the current directory**

```sh
fda myproject
```

This adds the following line to your shell profile (`.bashrc` or `.zshrc`):

```sh
alias go-myproject="cd /path/to/current/directory"
```

### **Adding an alias for a specific directory**

```sh
fda myproject /home/user/code/myproject
```

This adds:

```sh
alias go-myproject="cd /home/user/code/myproject"
```

### **Reloading your shell profile**

After adding an alias, reload your shell configuration:

```sh
source ~/.bashrc  # For bash users
source ~/.zshrc   # For zsh users
```

### **Navigating to an aliased directory**

Use the alias in your shell:

```sh
go-myproject
```

This will change the working directory to `/home/user/code/myproject`.

## **How It Works**

- The tool extracts the alias name and (optionally) a target directory from the command-line arguments.
- It validates the alias name and checks if it already exists in your shell profile.
- If valid, it appends the alias to your shell's configuration file (`.bashrc` or `.zshrc`).
- You can then use `source` to reload your shell profile and immediately use the alias.

## Usage

```
fda creates directory aliases for you to use

Usage:
  fda [command]

Available Commands:
  add
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for fda

Use "fda [command] --help" for more information about a command.
```
