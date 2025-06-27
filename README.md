```markdown
# green-cli

**Do you like green?**  
A simple CLI tool made with urfave and golang to backdate Git commits and colorize your GitHub contribution graph.  
Easily generate commits for today or simulate contribution streaks from the past!

---

## Features

- Create a commit for **today**  
- Backdate commits for **multiple days**  
- Print the **CLI version**  
- Automatically modifies a tracked file (`green_file.txt`) to trigger commits

---

## Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/green-cli.git
cd green-cli
```

### 2. Build the binary

```bash
go build -o green
```

This creates a `green` binary in your current directory.

**How to Install on Windows**  
**Option 1: Add to Environment Path**

1. Build the binary:

```bash
go build -o green.exe
```

2. Move `green.exe` to a directory (e.g., `C:\green-cli\`).

3. Add that directory to your Environment Variables:

- Open Start Menu → Search "Environment Variables"
- Click "Environment Variables"
- Under System variables → Find `Path` → Click Edit
- Click New → Add the path to `green.exe` (e.g., `C:\green-cli\`)
- Click OK and restart your terminal or system.

Now you can run `green` from anywhere!

**CLI Commands**

- `green today`  
  Create a Git commit for today (using current date).

- `green days --num=<N>`  
  Create backdated Git commits for the last N days.

- `green days --num=100`  
  Creates 100 daily commits starting from N days ago up to today.

- `green version`  
  Displays the current version of green-cli.

**Example Usage in Any Git Project**

```bash
# Navigate into your Git project
cd path/to/your/git/project

# Run a backdated commit
green days --num=30

# Or just commit for today
green today

#The tool will create (or modify) `green_file.txt` and commit it with the appropriate backdated metadata.
