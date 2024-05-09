Quad Checker
Overview
This repository contains a Go program called quadchecker that matches a given string to quad functions (quadA, quadB, quadC, quadD, quadE) and displays the name of the matching quad and its dimensions. If the argument is not a quad, the program prints "Not a quad function".

Usage
The program takes a string as an argument and determines if it matches any of the quad functions. If there is a match, it displays the name of the quad function along with its dimensions. If there are multiple matches, they are displayed alphabetically ordered and separated by ||.

Example Usage
css
Copy code
$ ./quadA 3 3 | go run .
[quadA] [3] [3]
css
Copy code
$ ./quadC 1 1 | go run .
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
bash
Copy code
$ echo -n "o--o"$'\n'"|"$'\n'"o" | go run .
Not a quad function
Contributing
Contributions and suggestions for improvements are welcome! Feel free to propose enhancements or report issues via pull requests or by opening an issue on GitHub.

License
This project is licensed under the MIT License. See the LICENSE file for details.

Feel free to adjust this description as needed for your repository.





