# Contributing to Forum

Thank you for considering contributing to Forum! We appreciate your time and effort and want to ensure that the contribution process is as smooth and transparent as possible.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
  - [Reporting Issues](#reporting-issues)
  - [Submitting Code Changes](#submitting-code-changes)
  - [Proposing New Features](#proposing-new-features)
- [Development Environment Setup](#development-environment-setup)
  - [BackEnd](#backend)
  - [FrontEnd](#frontend)
- [Testing](#testing)
- [Communication](#communication)

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [contact@example.com](mailto:contact@example.com).

## How to Contribute

### Reporting Issues

If you encounter any bugs, inconsistencies, or have suggestions for improvements, please open an issue on GitHub. When reporting an issue, please include:

- A clear and descriptive title.
- Steps to reproduce the issue.
- Any relevant screenshots or logs.
- The expected and actual behavior.

### Submitting Code Changes

1. **Fork the repository**: Create your own copy of the project.
2. **Create a new branch**: Use a descriptive name for your branch, e.g., `feature/add-authentication`.
3. **Make your changes**: Ensure your code adheres to the coding standards of the project.
4. **Test your changes**: Run the tests to verify that your changes don't break anything.
5. **Commit your changes**: Write a clear and concise commit message.
6. **Push to your branch**: Push the changes to your forked repository.
7. **Open a Pull Request**: Describe your changes and the problem they solve. Reference any related issues or PRs.

### Proposing New Features

If you have an idea for a new feature, please first open an issue to discuss the idea with the maintainers. This allows us to ensure that your work aligns with the project's goals before you invest significant time in it.

## Development Environment Setup

### BackEnd

The backend of our project is developed in Go (version 1.23.0). To set up your environment:

1. **Install Go**: Ensure Go 1.23.0 is installed on your system.
2. **Clone the repository**: `git clone https://github.com/longsizhuo/forum.git`
3. **Navigate to the BackEnd directory**: `cd BackEnd`
4. **Install dependencies**: Run `go mod tidy`.
5. **Generate gRPC code**:
   ```bash
   protoc --go_out=BackEnd\ --go-grpc_out=BackEnd\ BackEnd\proto\forum.proto
   protoc --go_out=BackEnd\ --go-grpc_out=BackEnd\ BackEnd\proto\chat.proto
   protoc --go_out=BackEnd\ --go-grpc_out=BackEnd\ BackEnd\proto\auth.proto
   ```
6. Build the project: Run go build ./....

### FrontEnd

The frontend is built using npm. To set up the frontend:

1. **Navigate to the FrontEnd directory**: `cd FrontEnd`
   
2. **Install dependencies**: Run npm install.
   
3. **Generate gRPC-Web code**:

```bash
protoc -I ./BackEnd/proto ./BackEnd/proto/auth.proto --js_out=import_style=commonjs:./FrontEnd/src/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./FrontEnd/src/proto
protoc -I ./BackEnd/proto ./BackEnd/proto/forum.proto --js_out=import_style=commonjs:./FrontEnd/src/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./FrontEnd/src/proto
protoc -I ./BackEnd/proto ./BackEnd/proto/chat.proto --js_out=import_style=commonjs:./FrontEnd/src/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./FrontEnd/src/proto
```

4. **Build the project**: Run npm run build.
   
### Testing
To ensure the stability and reliability of the project, we require that all changes be thoroughly tested.

- BackEnd Testing: Run go test ./... to execute backend tests.
- FrontEnd Testing: Run npm test to execute frontend tests.
  
Please write tests for any new functionality and ensure that all existing tests pass before submitting your code.

### Communication

For general questions or discussions, feel free to open a discussion on the GitHub repository or reach out via email at contact@longsizhuo.com

By contributing to this project, you agree to abide by our Code of Conduct. Thank you for your contributions!



