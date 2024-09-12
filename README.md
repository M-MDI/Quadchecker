# Quad Checker

## Overview

This repository contains a Go program called `quadchecker` that matches a given string to quad functions (`quadA`, `quadB`, `quadC`, `quadD`, `quadE`) and displays the name of the matching quad along with its dimensions. If the input doesn't match any quad function, the program prints "Not a quad function".

## Features

- Analyzes input strings to determine if they match any quad function patterns
- Supports multiple quad functions: quadA, quadB, quadC, quadD, quadE
- Displays matching quad function(s) with their dimensions
- Handles multiple matches, displaying them in alphabetical order

## Requirements

- Go 1.x or higher (replace x with the minimum version required)

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/quad-checker.git
   ```
2. Navigate to the project directory:
   ```
   cd quad-checker
   ```

## Usage

Run the program by piping the output of a quad function or a string to the `quadchecker`:

```
./quadA 3 3 | go run .
```

or

```
echo -n "o--o"$'\n'"|  |"$'\n'"o--o" | go run .
```

### Example Usage

1. Checking output from quadA:
   ```
   $ ./quadA 3 3 | go run .
   [quadA] [3] [3]
   ```

2. Checking output with multiple matches:
   ```
   $ ./quadC 1 1 | go run .
   [quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
   ```

3. Checking invalid input:
   ```
   $ echo -n "o--o"$'\n'"|"$'\n'"o" | go run .
   Not a quad function
   ```

## How It Works

1. The program reads the input string from standard input.
2. It analyzes the input string to determine its dimensions and pattern.
3. The program compares the input pattern against known quad function patterns.
4. If a match is found, it displays the quad function name(s) and dimensions.
5. If multiple matches are found, they are displayed in alphabetical order, separated by `||`.
6. If no match is found, it displays "Not a quad function".

## Project Structure

```
quad-checker/
├── main.go
├── quad_functions.go
├── pattern_matcher.go
├── README.md
└── LICENSE
```

## Contributing

Contributions are welcome! Here's how you can contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/AmazingFeature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
5. Push to the branch (`git push origin feature/AmazingFeature`)
6. Open a Pull Request

Please ensure your code adheres to the existing style and includes appropriate tests.

## Testing

To ensure the correctness of the quad checker, consider adding test cases for various input scenarios. You can use Go's built-in testing framework to create and run tests.

## Contact
Zangief
ig := Mehdim_dev
Discordid := 720655054834630676
Project Link: [https://github.com/yourusername/quad-checker](https://github.com/M-MDI/Quadchecker)
