<div align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" alt="Go Logo" width="400"/>
  
  # Go Mastery Path
  
  *A comprehensive, 150-program curriculum taking you from absolute Go beginner to advanced systems engineer.*
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
  [![Code Quality](https://img.shields.io/badge/lint-golangci--lint-blue)](https://golangci-lint.run/)
  [![License](https://img.shields.io/badge/license-MIT-green)](#)
</div>

---

This repository serves as a progressive reference and practical learning guide for the Go (Golang) programming language. The codebase is deliberately structured to scale in complexity, moving from basic syntax into standard library utilization, and finally into low-level primitives, metaprogramming, and high-throughput concurrency patterns.

## Curriculum & Repository Structure

The curriculum contains 150 complete, executable programs divided into three core progression levels.

### 1. Basics (`/01-basics`)
*Focuses on Go's core syntax, basic data structures, memory allocation, and standard execution flow.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Hello, World! | **26** | Basic Math Operations |
| **02** | Variables & Constants | **27** | Multi-Dimensional Slices |
| **03** | Control Flow (If/Else) | **28** | The `init()` Function |
| **04** | For Loops & Iteration | **29** | Custom Type Definitions |
| **05** | Arrays & Slices | **30** | Advanced Printf Formatting |
| **06** | Multiple Return Values | **31** | Bitwise Operations |
| **07** | Maps (Key-Value Pairs) | **32** | Variable Shadowing |
| **08** | Structs (Custom Types) | **33** | Slicing Slices (Re-slicing) |
| **09** | Pointers & Memory | **34** | Copying Slices (`copy`) |
| **10** | Simple Goroutines | **35** | Slice Length vs. Capacity |
| **11** | Switch Statements | **36** | Map Key Existence Check |
| **12** | The `defer` Keyword | **37** | Deleting from Maps (`delete`) |
| **13** | Variadic Functions | **38** | Labeled Loops (`break` / `continue`) |
| **14** | Type Conversion | **39** | Efficient String Concatenation |
| **15** | String Manipulation | **40** | Base64 Encoding & Decoding |
| **16** | Basic Error Creation | **41** | String Conversions (`strconv`) |
| **17** | Reading User Input | **42** | Raw String Literals |
| **18** | Struct Methods (Receivers) | **43** | Iterating Maps (Random Order) |
| **19** | Time & Date Formatting | **44** | Type Aliases vs Definitions |
| **20** | Random Number Generation | **45** | Anonymous Functions & Closures |
| **21** | Enums using `iota` | **46** | Struct Pointers |
| **22** | The Blank Identifier (`_`) | **47** | The Empty Interface |
| **23** | Panic and Recover | **48** | Typed vs Untyped Constants |
| **24** | Command-Line Arguments | **49** | Complex Numbers |
| **25** | Runes & String Iteration | **50** | The `goto` Statement |

### 2. Intermediate (`/02-intermediate`)
*Focuses on idiomatic Go, the standard library, interfaces, and foundational concurrency.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Interfaces & Polymorphism | **26** | Directory Traversal |
| **02** | Custom Error Handling | **27** | XML Marshalling |
| **03** | Channels & Synchronization | **28** | HTML Template Rendering |
| **04** | Worker Pool Pattern | **29** | HTTP Request Timeouts |
| **05** | Non-Blocking Select | **30** | Package Visibility Rules |
| **06** | Mutex & Thread-Safe State | **31** | Reading/Writing CSV Files |
| **07** | JSON Marshalling | **32** | URL Parsing |
| **08** | Basic HTTP Web Server | **33** | Hashing Data (SHA-256) |
| **09** | Context Cancellation | **34** | Custom String Formatting |
| **10** | File I/O Operations | **35** | Modifying Returns with `defer` |
| **11** | Struct Composition | **36** | Stateful HTTP Handlers |
| **12** | WaitGroups for Concurrency | **37** | Parsing Custom Time Strings |
| **13** | Type Assertions & Switches | **38** | Temporary Files & Directories |
| **14** | Custom Struct Sorting | **39** | Creating ZIP Archives |
| **15** | Command-Line Flags | **40** | HTTP Client with Custom Headers |
| **16** | HTTP Client Requests | **41** | Advanced JSON Struct Tags |
| **17** | Buffered Channels | **42** | Read-Write Mutex (`sync.RWMutex`) |
| **18** | Periodic Tasks (Tickers) | **43** | Advanced Error Handling |
| **19** | Environment Variables | **44** | JSON Stream Encoding/Decoding |
| **20** | Basic Unit Testing | **45** | Custom Sorting with `sort.Interface` |
| **21** | Single Execution (`sync.Once`) | **46** | In-Memory Data Streaming (`io.Pipe`) |
| **22** | Table-Driven Unit Tests | **47** | Timers and One-Shot Execution |
| **23** | Regular Expressions | **48** | HTTP Client with Retries |
| **24** | Executing OS Commands | **49** | SQL Database Connections |
| **25** | HTTP Panic Recovery | **50** | Method Overriding via Type Embedding |

### 3. Advanced (`/03-advanced`)
*Focuses on low-level primitives, system design patterns, metaprogramming, and performance optimization.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Worker Pool Graceful Shutdown | **26** | Memory Alignment & Padding |
| **02** | Object Pooling (`sync.Pool`) | **27** | Custom Reverse Proxy |
| **03** | Fan-In Concurrency Pattern | **28** | HTTP Transport RoundTripper |
| **04** | Lock-Free Atomic Operations | **29** | HTTP Request Tracing |
| **05** | Custom HTTP Middleware | **30** | Context Post-Cancellation |
| **06** | Reflection & Struct Tags | **31** | Functional Generics |
| **07** | Zero-Copy Memory Conversion | **32** | Fuzz Testing |
| **08** | Pipeline & Error Propagation | **33** | Raw TCP Layer 4 Proxy |
| **09** | Token Bucket Rate Limiting | **34** | Custom Context Implementation |
| **10** | OS Signal Handling | **35** | Zero-Allocation JSON Stream Parsing |
| **11** | Request-Scoped Context Data | **36** | Unsafe Pointer Arithmetic |
| **12** | Generics (Type-Safe Data) | **37** | Reusable Byte Buffer Pool |
| **13** | Functional Options Pattern | **38** | Exclusive OS File Locking |
| **14** | Bounded Concurrency (Semaphores)| **39** | Custom Iterator Pattern |
| **15** | Custom JSON Unmarshaling | **40** | Leaky Bucket Rate Limiter |
| **16** | Dynamic Struct Mutation | **41** | Dynamic Code Loading (`plugin`) |
| **17** | In-Memory Pub/Sub Event Bus | **42** | Detached Contexts |
| **18** | Runtime Memory Profiling | **43** | Memory-Mapped Files (`Mmap`) |
| **19** | TCP Server with Deadlines | **44** | Generic LRU Cache |
| **20** | Singleflight Request Coalescing | **45** | Socket Options (`SO_REUSEPORT`) |
| **21** | Calling C Code (CGO) | **46** | Dynamic Function Generation |
| **22** | Concurrent Cache (`sync.Map`) | **47** | Capturing Call Stack Traces |
| **23** | Packing Binary Protocols | **48** | AST Modification & Rewriting |
| **24** | AST Parsing & Metaprogramming | **49** | Priority Queue (`container/heap`) |
| **25** | Broadcast Signaling (`sync.Cond`) | **50** | WebAssembly Integration |

---

## Getting Started

### Prerequisites
- Go 1.21 or higher (required for Generics and modern Context features in the advanced section).

### Installation & Execution

1. Clone the repository:
   ```bash
   git clone [https://github.com/rohitmohan/golang-mastery-path.git](https://github.com/rohitmohan/golang-mastery-path.git)
   cd golang-mastery-path