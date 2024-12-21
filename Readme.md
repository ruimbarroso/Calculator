# Go Calculator with Fyne GUI

This project is a calculator application built using the Go programming language and the Fyne GUI library. The calculator supports a variety of mathematical operations and provides a simple graphical user interface (GUI) for interacting with it. Additionally, it features a history log that keeps track of previously executed equations.

![[Pasted image 20241219180633.png]]
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
- Go (version 1.23.1 or above)
- Fyne (GUI library for Go, v2.5.2)

### Steps to Install:

1. Clone the repository to your local machine:
```bash
    git clone https://github.com/yourusername/yourrepositoryname.git
````
    
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
    

### Running the Application

Once the application is running, the GUI will open, allowing you to perform calculations and view the history of your equations.

## How It Works

The application follows the standard order of operations (PEMDAS):

1. **Parentheses**
2. **Exponentiation**
3. **Multiplication/Division**
4. **Addition/Subtraction**

The core logic for evaluating mathematical expressions is implemented in the Go code, and the results are displayed via the Fyne GUI. The user can enter expressions into the input field, and the result is computed and displayed immediately.

## Example Operations:

- `2 + 2` will give the result: `4`
- `3 * (2 + 4)` will compute `3 * 6` and give the result: `18`
- `2 ^ 3` will compute `2^3` and give the result: `8`
- `100 l 10` will compute `log(100) base 10` and give the result: `2`

## License

This project is licensed under the MIT License - see the LICENSE file for details.