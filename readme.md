<div align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg" alt="Go Logo" width="400"/>
  
  # Go Mastery Path
  
  *A comprehensive, 60-program curriculum taking you from absolute Go beginner to advanced systems engineer.*
  
  [![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)](https://go.dev/)
  [![Code Quality](https://img.shields.io/badge/lint-golangci--lint-blue)](https://golangci-lint.run/)
  [![License](https://img.shields.io/badge/license-MIT-green)](#)
</div>

---

This repository serves as a progressive reference and practical learning guide for the Go (Golang) programming language. The codebase is deliberately structured to scale in complexity, moving from basic syntax into standard library utilization, and finally into low-level primitives and high-throughput concurrency patterns.

## 📁 Curriculum & Repository Structure

The curriculum contains 60 complete, executable programs divided into three core progression levels.

### 1. Basics (`/01-basics`)
*Focuses on Go's core syntax, basic data structures, memory allocation, and standard execution flow.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Hello, World! | **11** | Switch Statements |
| **02** | Variables & Constants | **12** | The `defer` Keyword |
| **03** | Control Flow (If/Else) | **13** | Variadic Functions |
| **04** | For Loops & Iteration | **14** | Type Conversion |
| **05** | Arrays & Slices | **15** | String Manipulation |
| **06** | Multiple Return Values | **16** | Basic Error Creation |
| **07** | Maps (Key-Value Pairs) | **17** | Reading User Input |
| **08** | Structs (Custom Types) | **18** | Struct Methods (Receivers) |
| **09** | Pointers & Memory | **19** | Time & Date Formatting |
| **10** | Simple Goroutines | **20** | Random Number Generation |

### 2. Intermediate (`/02-intermediate`)
*Focuses on idiomatic Go, the standard library, interfaces, and foundational concurrency.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Interfaces & Polymorphism | **11** | Struct Composition (Embedding) |
| **02** | Custom Error Handling | **12** | WaitGroups for Concurrency |
| **03** | Channels & Synchronization | **13** | Type Assertions & Switches |
| **04** | Worker Pool Pattern | **14** | Custom Struct Sorting |
| **05** | Non-Blocking Select & Timeouts| **15** | Command-Line Flags |
| **06** | Mutex & Thread-Safe State | **16** | HTTP Client Requests |
| **07** | JSON Marshalling | **17** | Buffered Channels |
| **08** | Basic HTTP Web Server | **18** | Periodic Tasks (Tickers) |
| **09** | Context Cancellation | **19** | Environment Variables |
| **10** | File I/O Operations | **20** | Basic Unit Testing |

### 3. Advanced (`/03-advanced`)
*Focuses on low-level primitives, system design patterns, metaprogramming, and performance optimization.*

| # | Topic / Program | # | Topic / Program |
| :--- | :--- | :--- | :--- |
| **01** | Graceful Shutdown & Context | **11** | Request-Scoped Context Data |
| **02** | Object Pooling (`sync.Pool`) | **12** | Generics (Type-Safe Data) |
| **03** | Fan-In Concurrency Pattern | **13** | Functional Options Pattern |
| **04** | Lock-Free Atomic Operations | **14** | Bounded Concurrency (Semaphores) |
| **05** | Custom HTTP Middleware | **15** | Custom JSON Unmarshaling |
| **06** | Reflection & Struct Tags | **16** | Dynamic Struct Mutation |
| **07** | Zero-Copy Memory Conversion | **17** | In-Memory Pub/Sub Event Bus |
| **08** | Pipeline & Error Propagation | **18** | Runtime Memory Profiling |
| **09** | Token Bucket Rate Limiting | **19** | TCP Server with Deadlines |
| **10** | OS Signal Handling | **20** | Singleflight Request Coalescing |

---

##  Getting Started

### Prerequisites
* Go 1.18 or higher (required for Generics in the advanced section).

### Installation & Execution

1. Clone the repository:
   ```bash
   git clone [https://github.com/rohitmohan/golang-mastery-path.git](https://github.com/rohitmohan/golang-mastery-path.git)
   cd golang-mastery-path