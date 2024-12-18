# Electronic Sign Application

### Intro
Imagine you have a digital bulletin board, like the ones you might see at a train station or sports stadium.

This board is like a grid of tiny light bulbs - 6 rows (labeled A to F) and 36 columns (numbered 0 to 35). Each bulb can be either on or off.

The program works like this:
Think of it like giving directions to turn on specific lights:
When you type something like "A1", you're saying "turn on the light in row A, column 1"

The program does three main things:
1. Takes input instructions (like "A2B12C34")
2. Figures out which lights to turn on (like a translator)
3. Shows you the result using stars (*) for lit bulbs



 **Table of content:**
 - [Challenges](#item1)
 - [Usage](#item2)
 - [Installation](#item3)
    - [Build](#item4)
    - [Docker](#item5)
    - [Test](#item6)
 - [Implementation](#item6.1)
    - [Data Structure](#item7)
    - [Flow](#item8)
    - [Entities](#item9)
    - [Transformation](#item10)
    - [Storage](#item11)
    - [View](#item12)
    - [Characters](#item13)
    - [Display](#item14)
    - [Memory Clean Up](#item15)
 - [Project Structure](#item16)
 - [Fault Tolerance](#item17)

 <!-- headings -->
 <a id="item1"></a>
### Challenges
- Maintaining modularity for future extensions
- System design principles
- Bitmap display systems
- Memory management by using Flyweight pattern for reserved characters
- TDD

 <!-- headings -->
 <a id="item2"></a>
 

### Usage
```
Electronic Sign CLI
Commands:
  add <type> <text> - Add a new view (type: pixel or character)
  show            - Display all views
  clear           - Clear all views
  exit            - Exit the program
```
>>>

> add pixel E35B12C5

> add character ABC123

>>> View added with character transformer

> show

>>> Displaying all views:

 <!-- headings -->
 <a id="item3"></a>
 

### Installation Guide

+ Manual Installation (Go)
    - Install Go (version 1.21 or later)
    - Docker 
    - Clone the repository
    - Build and run the application

 <!-- headings -->
 <a id="item4"></a>
 

### Build and run the application

```    
    go build -o elecsign main.go
    
    ./elecsign
```
 <!-- headings -->
 <a id="item5"></a>


### Docker Installation
```
    docker build -t elecsign .
    docker run -it elecsign
    docker run -it --rm elecsign
```
 <!-- headings -->
 <a id="item6"></a>

### Test
```
    go test -v ./...
```


 <!-- headings -->
 <a id="item6.1"></a>
## Implementation


* Enter a view as a sequence of pixels and save it in memory
* Print all views stored in memory
* Clear the memory

 <!-- headings -->
 <a id="item7"></a>

### Grid Data Structure
As the size of array for saving the bitwise (On/Off) is fixed it would be more efficient to take it a fixed array size so, size is known at compile time, allowing for better memory allocation
- Memory is allocated on the stack rather than the heap, reducing garbage collection overhead.
- Contiguous memory layout improves cache locality and performance


###### Why bytearray?
Each pixel needs only 1 bit (on/off)
216(6*36) pixels can be stored in 27 bytes (216/8 = 27)
Allows bitwise operations for efficient manipulation

```
Using bool: 8 bytes (1 byte per pixel) - unnecessary memory overhead
Using byte: 1 byte (1 bit per pixel) - fit within the size of grid
```
 <!-- headings -->
 <a id="item8"></a>

### Data Flow
```
    User input → CLI Controller
    Controller → Transformer
    Transformer → View
    View → Display
    Display → Renderer
    Renderer → Console Output
```
 <!-- headings -->
 <a id="item9"></a>

### Core Entities
1. <b>Display</b>: Manages the visual representation (it only supports console interface)
2. <b>Grid</b>: Represents the physical structure of the sign in memory
3. <b>View</b>: Represents a single displayable pattern
4. <b>Transformer</b>: Converts input into displayable patterns (indexes of bytes and bits)

 <!-- headings -->
 <a id="item10"></a>

##### String to Coordinate Transformation

The algorithm converts user input strings representing pixel coordinates (e.g., "A0B12C5") into memory-mapped coordinates for a 6x36 pixel display, achieving O(n) time complexity. it works with two pointer: first pointer start with the first char and the second pointer move until the next character, then first pointer move to the current index of the array and repeat this unit end of the array

```
// Input: "A0B12C5"
// Iteration states:
// i=1: currentChar='0', segment="A"
// i=2: currentChar='B', segment="A0" -> Coordinate{row: 0, col: 0}
// i=4: currentChar='2', segment="B1"
// i=5: currentChar='C', segment="B12" -> Coordinate{row: 1, col: 12}
```
 <!-- headings -->
 <a id="item11"></a>

##### Bit-Level Storage

The implementation uses bit-level operations to efficiently store and retrieve pixel states in a byte array, where each pixel's state (on/off) is represented by a single bit.
```
// For coordinate (2,5):
// position = 2 * 36 + 5 = 77
// byteIndex = 77 / 8 = 9
// bitIndex = 77 % 8 = 5
// 
// Byte at index 9:  0 0 1 0 0 1 0 0 -- position 3 and 6 were set before
// Mask (bit 5):     0 0 0 0 1 0 0 0 -- position 5
// Result:           0 0 1 0 1 1 0 0 -- now position 5, 3, 6 is set
```
 <!-- headings -->
 <a id="item12"></a>

##### View

The View implementation entity fo keeping all user input coordinates and to represent them
```
// Memory representation for "A0B0F35":
// View 1:
// Byte 0: 10000000 (A0)
// Byte 1: 00000000
// Byte 4: 00000100 (B0)
// ... remaining bytes
// Byte 26: 00000001 (F35)
```
 <!-- headings -->
 <a id="item13"></a>

##### Character Transformer

The Character Transformer is a specialized implementation of the Transformer interface that handles predefined character patterns (A-Z, 0-9) by maintaining a mapping of characters to their pixel coordinates.

```
// Example pattern for letter 'A':
// It contains the coordinations of the pixels in the grid
// byte index: 0, bit index: 2 and ...
// 02,03,11,14,20,21,22,23,24,30,34,40,44,50,54,60,61,62,63,64
//
// Visual representation:
//   * *    
//  *   *   
// *     *  
// * * * * 
// *     *  
// *     *  
```

 <!-- headings -->
 <a id="item14"></a>

##### Display
The Console Display entity is responsible for rendering the 6x36 pixel grid as ASCII in the terminal, where (*) represents ON pixels and spaces represent OFF pixels.

Instead of creating new strings for each modification, ```strings.Builder``` in go maintains a single growing buffer. By pre-allocating the exact space needed (row * (column + 1)), it avoids multiple reallocations during the rendering process.
``` 
The Renderable interface allows our application to support multiple display types (console, LED matrix, LCD, etc.) through a common rendering contract.
```
 <!-- headings -->
 <a id="item15"></a>

### Memory Clean Up
The display manager cleanup the memory after showing the items (show command). To avoid constantly re allocating memory
```
// Clear all views
func (v *View) Clear() {
    vm.views = vm.views[:0]  // Reuse underlying array
}
```
 <!-- headings -->
 <a id="item16"></a>

### Project Structure 

```
electronic-sign/
├── cmd/
│   └── cli.go
├── internal/
│   ├── display/
│   ├── grid/
│   ├── transformer/
│   └── view/
│       └── view.go
├── Dockerfile
└── docker-compose.yml
```
 <!-- headings -->
 <a id="item17"></a>

### Fault Tolerance
If you make a mistake (like typing "G1" when there's no row G), the program simply ignores it and continues with the valid instructions.
