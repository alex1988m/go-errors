# Error Handling in Go

This repository demonstrates different approaches to error handling in Go, specifically focusing on matching errors using static and dynamic error messages, as well as the use of custom error types.

## Table of Contents
- [Error Handling in Go](#error-handling-in-go)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [File Descriptions](#file-descriptions)
  - [Usage](#usage)
    - [Listings](#listings)
  - [Error Handling Guidance](#error-handling-guidance)
  - [Running the Code](#running-the-code)

## Overview

This codebase contains examples of different error handling techniques in Go. The main goal is to illustrate how to:
- Use `errors.New` for static error messages.
- Use `fmt.Errorf` for dynamic error messages.
- Match errors using top-level variables.
- Implement custom error types for more complex error matching.

## File Descriptions

- **main.go**: Contains the `main` function and the various listings demonstrating different error handling scenarios.
- **user.go**: Defines the `User` struct and the `GetUser` function, which simulates fetching a user and returning different types of errors based on the parameters.
- **errors.go**: Defines the `ErrNotFound` variable and the `UserNotFoundError` custom error type. Also includes the `logError` function to handle logging of different errors.

## Usage

The code is structured into five listings, each demonstrating a unique error handling scenario. The `main` function runs all five listings in sequence.

### Listings

1. **Listing 1**: Demonstrates handling a static error that is not matched.
2. **Listing 2**: Demonstrates handling a dynamic error that is not matched.
3. **Listing 3**: Demonstrates handling a static error that is matched using a top-level variable.
4. **Listing 4**: Demonstrates handling a dynamic error that is matched using a custom error type.
5. **Listing 5**: Demonstrates handling wrapped errors, showing the propagation of errors up the call stack.

## Error Handling Guidance

| Error Matching? | Error Message | Guidance                                   |
|-----------------|---------------|--------------------------------------------|
| No              | Static        | Use `errors.New`                           |
| No              | Dynamic       | Use `fmt.Errorf`                           |
| Yes             | Static        | Use top-level variable with `errors.New`   |
| Yes             | Dynamic       | Use custom error type                      |

## Running the Code

To run the code, follow these steps:

1. **Clone the repository**:
   ```sh
   git clone https://github.com/alex1988m/go-errors.git
   cd go-errors
