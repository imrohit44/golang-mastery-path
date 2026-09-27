<div align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" alt="Go Logo" width="400"/>
  
  # Go Mastery Path
  
  *A comprehensive, 120-program curriculum taking you from absolute Go beginner to advanced systems engineer.*
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
  [![Code Quality](https://img.shields.io/badge/lint-golangci--lint-blue)](https://golangci-lint.run/)
  [![License](https://img.shields.io/badge/license-MIT-green)](#)
</div>

---

This repository serves as a progressive reference and practical learning guide for the Go (Golang) programming language. The codebase is deliberately structured to scale in complexity, moving from basic syntax into standard library utilization, and finally into low-level primitives, metaprogramming, and high-throughput concurrency patterns.

##  Curriculum & Repository Structure

The curriculum contains 120 complete, executable programs divided into three core progression levels.

### 1. Basics (`/01-basics`)
*Focuses on Go's core syntax, basic data structures, memory allocation, and standard execution flow.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Hello, World! | **21** | Enums using `iota` |
| **02** | Variables & Constants | **22** | The Blank Identifier (`_`) |
| **03** | Control Flow (If/Else) | **23** | Panic and Recover |
| **04** | For Loops & Iteration | **24** | Command-Line Arguments |
| **05** | Arrays & Slices | **25** | Runes & String Iteration |
| **06** | Multiple Return Values | **26** | Basic Math Operations |
| **07** | Maps (Key-Value Pairs) | **27** | Multi-Dimensional Slices |
| **08** | Structs (Custom Types) | **28** | The `init()` Function |
| **09** | Pointers & Memory | **29** | Custom Type Definitions |
| **10** | Simple Goroutines | **30** | Advanced Printf Formatting |
| **11** | Switch Statements | **31** | Bitwise Operations |
| **12** | The `defer` Keyword | **32** | Variable Shadowing |
| **13** | Variadic Functions | **33** | Slicing Slices (Re-slicing) |
| **14** | Type Conversion | **34** | Copying Slices (`copy`) |
| **15** | String Manipulation | **35** | Slice Length vs. Capacity |
| **16** | Basic Error Creation | **36** | Map Key Existence Check |
| **17** | Reading User Input | **37** | Deleting from Maps (`delete`) |
| **18** | Struct Methods (Receivers) | **38** | Labeled Loops (`break` / `continue`) |
| **19** | Time & Date Formatting | **39** | Efficient String Concatenation |
| **20** | Random Number Generation | **40** | Base64 Encoding & Decoding |

### 2. Intermediate (`/02-intermediate`)
*Focuses on idiomatic Go, the standard library, interfaces, and foundational concurrency.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Interfaces & Polymorphism | **21** | Single Execution (`sync.Once`) |
| **02** | Custom Error Handling | **22** | Table-Driven Unit Tests |
| **03** | Channels & Synchronization | **23** | Regular Expressions |
| **04** | Worker Pool Pattern | **24** | Executing OS Commands |
| **05** | Non-Blocking Select | **25** | HTTP Panic Recovery |
| **06** | Mutex & Thread-Safe State | **26** | Directory Traversal |
| **07** | JSON Marshalling | **27** | XML Marshalling |
| **08** | Basic HTTP Web Server | **28** | HTML Template Rendering |
| **09** | Context Cancellation | **29** | HTTP Request Timeouts |
| **10** | File I/O Operations | **30** | Package Visibility Rules |
| **11** | Struct Composition | **31** | Reading/Writing CSV Files |
| **12** | WaitGroups for Concurrency | **32** | URL Parsing |
| **13** | Type Assertions & Switches | **33** | Hashing Data (SHA-256) |
| **14** | Custom Struct Sorting | **34** | Custom String Formatting |
| **15** | Command-Line Flags | **35** | Modifying Returns with `defer` |
| **16** | HTTP Client Requests | **36** | Stateful HTTP Handlers |
| **17** | Buffered Channels | **37** | Parsing Custom Time Strings |
| **18** | Periodic Tasks (Tickers) | **38** | Temporary Files & Directories |
| **19** | Environment Variables | **39** | Creating ZIP Archives |
| **20** | Basic Unit Testing | **40** | HTTP Client with Custom Headers |

### 3. Advanced (`/03-advanced`)
*Focuses on low-level primitives, system design patterns, metaprogramming, and performance optimization.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Worker Pool Graceful Shutdown | **21** | Calling C Code (CGO) |
| **02** | Object Pooling (`sync.Pool`) | **22** | Concurrent Cache (`sync.Map`) |
| **03** | Fan-In Concurrency Pattern | **23** | Packing Binary Protocols |
| **04** | Lock-Free Atomic Operations | **24** | AST Parsing & Metaprogramming |
| **05** | Custom HTTP Middleware | **25** | Broadcast Signaling (`sync.Cond`) |
| **06** | Reflection & Struct Tags | **26** | Memory Alignment & Padding |
| **07** | Zero-Copy Memory Conversion | **27** | Custom Reverse Proxy |
| **08** | Pipeline & Error Propagation | **28** | HTTP Transport RoundTripper |
| **09** | Token Bucket Rate Limiting | **29** | HTTP Request Tracing |
| **10** | OS Signal Handling | **30** | Context Post-Cancellation |
| **11** | Request-Scoped Context Data | **31** | Functional Generics |
| **12** | Generics (Type-Safe Data) | **32** | Fuzz Testing |
| **13** | Functional Options Pattern | **33** | Raw TCP Layer 4 Proxy |
| **14** | Bounded Concurrency (Semaphores)| **34** | Custom Context Implementation |
| **15** | Custom JSON Unmarshaling | **35** | Zero-Allocation JSON Stream Parsing |
| **16** | Dynamic Struct Mutation | **36** | Unsafe Pointer Arithmetic |
| **17** | In-Memory Pub/Sub Event Bus | **37** | Reusable Byte Buffer Pool |
| **18** | Runtime Memory Profiling | **38** | Exclusive OS File Locking |
| **19** | TCP Server with Deadlines | **39** | Custom Iterator Pattern |
| **20** | Singleflight Request Coalescing | **40** | Leaky Bucket Rate Limiter |

---

##  Getting Started

### Prerequisites
- Go 1.21 or higher (required for Generics and modern Context features in the advanced section).

### Installation & Execution

1. Clone the repository:
   ```bash
   git clone [https://github.com/rohitmohan/golang-mastery-path.git](https://github.com/rohitmohan/golang-mastery-path.git)
   cd golang-mastery-path