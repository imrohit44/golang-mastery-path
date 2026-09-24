<div align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" alt="Go Logo" width="400"/>
  
  # Go Mastery Path
  
  *A comprehensive, 90-program curriculum taking you from absolute Go beginner to advanced systems engineer.*
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
  [![Code Quality](https://img.shields.io/badge/lint-golangci--lint-blue)](https://golangci-lint.run/)
  [![License](https://img.shields.io/badge/license-MIT-green)](#)
</div>

---

This repository serves as a progressive reference and practical learning guide for the Go (Golang) programming language. The codebase is deliberately structured to scale in complexity, moving from basic syntax into standard library utilization, and finally into low-level primitives, metaprogramming, and high-throughput concurrency patterns.

## 📁 Curriculum & Repository Structure

The curriculum contains 90 complete, executable programs divided into three core progression levels.

### 1. Basics (`/01-basics`)
*Focuses on Go's core syntax, basic data structures, memory allocation, and standard execution flow.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Hello, World! | **16** | Basic Error Creation |
| **02** | Variables & Constants | **17** | Reading User Input |
| **03** | Control Flow (If/Else) | **18** | Struct Methods (Receivers) |
| **04** | For Loops & Iteration | **19** | Time & Date Formatting |
| **05** | Arrays & Slices | **20** | Random Number Generation |
| **06** | Multiple Return Values | **21** | Enums using `iota` |
| **07** | Maps (Key-Value Pairs) | **22** | The Blank Identifier (`_`) |
| **08** | Structs (Custom Types) | **23** | Panic and Recover |
| **09** | Pointers & Memory | **24** | Command-Line Arguments |
| **10** | Simple Goroutines | **25** | Runes & String Iteration |
| **11** | Switch Statements | **26** | Basic Math Operations |
| **12** | The `defer` Keyword | **27** | Multi-Dimensional Slices |
| **13** | Variadic Functions | **28** | The `init()` Function |
| **14** | Type Conversion | **29** | Custom Type Definitions |
| **15** | String Manipulation | **30** | Advanced Printf Formatting |

### 2. Intermediate (`/02-intermediate`)
*Focuses on idiomatic Go, the standard library, interfaces, and foundational concurrency.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Interfaces & Polymorphism | **16** | HTTP Client Requests |
| **02** | Custom Error Handling | **17** | Buffered Channels |
| **03** | Channels & Synchronization | **18** | Periodic Tasks (Tickers) |
| **04** | Worker Pool Pattern | **19** | Environment Variables |
| **05** | Non-Blocking Select | **20** | Basic Unit Testing |
| **06** | Mutex & Thread-Safe State | **21** | Single Execution (`sync.Once`) |
| **07** | JSON Marshalling | **22** | Table-Driven Unit Tests |
| **08** | Basic HTTP Web Server | **23** | Regular Expressions |
| **09** | Context Cancellation | **24** | Executing OS Commands |
| **10** | File I/O Operations | **25** | HTTP Panic Recovery |
| **11** | Struct Composition | **26** | Directory Traversal |
| **12** | WaitGroups for Concurrency | **27** | XML Marshalling |
| **13** | Type Assertions & Switches | **28** | HTML Template Rendering |
| **14** | Custom Struct Sorting | **29** | HTTP Request Timeouts |
| **15** | Command-Line Flags | **30** | Package Visibility Rules |

### 3. Advanced (`/03-advanced`)
*Focuses on low-level primitives, system design patterns, metaprogramming, and performance optimization.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Worker Pool Graceful Shutdown | **16** | Dynamic Struct Mutation |
| **02** | Object Pooling (`sync.Pool`) | **17** | In-Memory Pub/Sub Event Bus |
| **03** | Fan-In Concurrency Pattern | **18** | Runtime Memory Profiling |
| **04** | Lock-Free Atomic Operations | **19** | TCP Server with Deadlines |
| **05** | Custom HTTP Middleware | **20** | Singleflight Request Coalescing |
| **06** | Reflection & Struct Tags | **21** | Calling C Code (CGO) |
| **07** | Zero-Copy Memory Conversion | **22** | Concurrent Cache (`sync.Map`) |
| **08** | Pipeline & Error Propagation | **23** | Packing Binary Protocols |
| **09** | Token Bucket Rate Limiting | **24** | AST Parsing & Metaprogramming |
| **10** | OS Signal Handling | **25** | Broadcast Signaling (`sync.Cond`) |
| **11** | Request-Scoped Context Data | **26** | Memory Alignment & Padding |
| **12** | Generics (Type-Safe Data) | **27** | Custom Reverse Proxy |
| **13** | Functional Options Pattern | **28** | HTTP Transport RoundTripper |
| **14** | Bounded Concurrency (Semaphores)| **29** | HTTP Request Tracing |
| **15** | Custom JSON Unmarshaling | **30** | Context Post-Cancellation |

---

## Getting Started

### Prerequisites
* Go 1.21 or higher (required for Generics and modern Context features in the advanced section).

### Installation & Execution

1. Clone the repository:
   ```bash
   git clone [https://github.com/rohitmohan/golang-mastery-path.git](https://github.com/rohitmohan/golang-mastery-path.git)
   cd golang-mastery-path