# Trivil Programming Language

This repository contains the implementation of a code generator targeting the JVM for Trivil programming language, developed as part of a university thesis.

## Repository Contents

- **`src/`**: Contains the implementation of the Trivil compiler with the JVM backend extension.
- **`run.sh`**: Automates the program execution, starting from running the compiler to generate bytecodes and executing Java classes.
- **`examples/`**: Includes code examples written in Trivil, showcasing different features along with the expected outputs.
- **`runtime/`**: Contains predefined Jasmin files such as standard input and output, which can be used for function modifications.
- **`jasmin-2.4/`**: Jasmin assembler for compiling bytecodes in plain text into Java classes.
- **`vscode-extension/`**: Includes a VSCode extension for the Trivil programming language.
- **`trivil-v0.72/`**: The current Trivil compiler, which targets the C programming language.
- **`стд/`**: Trivil standard library. Note that not all features are supported by the JVM backend at this time.
- **`config/`**: Contains files such as error descriptions used by the compiler and build commands for building the executable for the C language.
