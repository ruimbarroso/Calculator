# Go Calculator with Fyne GUI

This project is a calculator application built using the Go programming language and the Fyne GUI library. The calculator supports a variety of mathematical operations and provides a simple graphical user interface (GUI) for interacting with it. Additionally, it features a history log that keeps track of previously executed equations.

![Calculator App image](https://github.com/ruimbarroso/Calculator/blob/main/images/Screenshot%202024-12-19%20180615.png)

## Features

### Supported Operations:
The calculator supports the following mathematical operations:

- **Sum (`+`)**
- **Subtraction (`-`)**
- **Multiplication (`*`)**
- **Division (`/`)**
- **Exponentiation (`^`)**
- **Root (`r`)**
- **Logarithm (`l`)**
- **Parentheses (`()`)** for grouping operations

### History:
The application keeps a history of all the equations that have been executed. This allows you to review previous calculations without needing to re-enter them.

## Installation

### Prerequisites:
- Go (version 1.23.1 or later)
- Fyne (fyne.io/fyne/v2, version 2.5.4) – GUI library for Go
- Litter (github.com/sanity-io/litter, version 1.5.8) – Library for pretty-printing Go data structures

### Steps to Install:

1. Clone the repository to your local machine:
```bash
    git clone https://github.com/yourusername/yourrepositoryname.git
```
    
2. Navigate to the project directory:
```bash
cd yourrepositoryname
```
    
3. Install the required dependencies:
```bash
go get fyne.io/fyne/v2
```
    
4. Build and run the application:
```bash
go run main.go
```

## Pre-built Binary for Windows

For convenience, a pre-built executable for Windows 11/10 (64-bit) is already available in the build folder. You can run the calculator.exe directly without the need for compilation.

### Running the Application

Once the application is running, the GUI will open, allowing you to perform calculations and view the history of your equations.

## How It Works

The application follows the standard order of operations (PEMDAS):

1. **Parentheses**
2. **Exponentiation**
3. **Multiplication/Division**
4. **Addition/Subtraction**

The core logic for evaluating mathematical expressions is implemented in the Go code, and the results are displayed via the Fyne GUI. The user can enter expressions into the input field, and the result is computed and displayed immediately.

The parsing of mathematical expressions is achieved using a Pratt Parser approach with lookup tables. This design allows for a flexible and efficient handling of operator precedence and associativity, making it easier to extend and maintain the parsing logic.

## Example Operations:

- `2 + 2` will give the result: `4`
- `3 * (2 + 4)` will compute `3 * 6` and give the result: `18`
- `2 ^3` will compute `2^3` and give the result: `8`
- `8 r3` will compute `8^(1/3)` and give the result: `2`
- `100 l10` will compute `log(100) base 10` and give the result: `2`

## License

This project is licensed under the MIT License - see the LICENSE file for details.