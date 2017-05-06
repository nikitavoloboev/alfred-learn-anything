# [1952: compilers ✨](https://my.mindnode.com/JxLJTw9UYqexCqyaqdKPxGGioWsDsWDmASWxpfsp)

- is a computer program (or a set of programs) that transforms source code written in a programming language (the source language) into another computer language (the target language), with the latter often having a binary form known as object code  
	- is often done to create an executable program


## [JIT](https://my.mindnode.com/LgpsEFf3gzioSzG7bFNbMRcRMxQgjB8EtJDAD7u4)

- is compilation done during execution of a program (at run time), rather than prior to execution  
- mist often this consists of translation to machine code, which is then executed directly, but can also refer to translation to another format  
- a system implementing a JIT compiler typically continuously analyses the code being executed and identifies parts of the code where the speedup gained from compilation would outweigh the overhead of compiled that code  
- JIT compilation is a combination of two traditional approaches to translation to machine code  
	1. ahead of time compilation (AOT)  
	2. and interpretation

## [interpreter](https://my.mindnode.com/pezFduzLe1ot7ZKv5qckzijpwDQFDZrNs1cgjWhs)

- is a computer program that directly executes (performs) instructions written in a programming or scripting language, without previously compiling them into machine language program  
- an interpreter generally uses one of the following strategies for program execution :   
	1. parse the source code and perform its behaviour directly (lisp, BASIC)  
	2. translate source code into some efficient intermediate representation and immediately execute this (perl, python, ruby)  
	3. explicitly execute stored precompiled code made by a compiler which is part of the interpreter system (Pascal)  
- source programs are compiled ahead of time and stored as machine independent code, which is then linked at run-time and executed by an interpreter and/or compiler (for JIT systems)

## [bytecode](https://my.mindnode.com/RPbjdQipHWeALK23FTASbKF7ohxd7tNECm3Ntwkt)

- is a form of instruction set designed for efficient execution by a software interpreter  
- unlike human-readable source code, byte codes are compact numeric codes, constants and references (normally numeric addresses) that encode the result of compiler parsing and semantic analysis of things like type and nesting depths of program objects  
- they thus allow much better performance than interpreting source code directly  
- authors of V8 and Dart have challenged the notion that intermediate byte code is needed for fast and efficient VM implementation, both of these languages do direct JIT compiling from source code to machine code with no byte code intermediary

## [object code](https://my.mindnode.com/wNkD1fSF9u7EpYKpWDwSFzNSZfiLGSSuowQtsw7r)

- is what compiler produces   
- it is a sequence of statements or instructions in a computer language, usually machine code language (i.e. binary) or an intermediate language such as register transfer language (RTL)  
- object files can in turn be linked to form an executable file or library file  
	- in order to be used, object code must either be placed in an executable file, a library file, or an object file  
- object code is a portion of machine code that has not yet been linked into a complete program

## [machine code](https://my.mindnode.com/fixbcG2yGbd5D5Xqs6hXQeH9zGxZqHxaA7uHmfpq)

- is a set of instructions executed directly by a computer’s central processing unit (CPU)  
- each instruction performs a very specific task, such as a load, a jump or an ALU operation on a unit of data in a CPU register or memory  
	- every program directly executed by a CPU is made up of a series of such instructions

## [compiler design](https://my.mindnode.com/zi5qk8yjNq6u67xzqcCapN9QGqs8aeyTqLTqtJw9)

- how to create an efficient and optimised compiler

## [optimisation](https://my.mindnode.com/Gp7bBEjqcvn3RzzyUgRqPTwXwgTA6o9bx6yvXzpa)

- how to effectively optimise the process of compiling

- [optimising compiler](https://my.mindnode.com/BgnCykCteFedakmQYGqqJnZS1SFufzAMUUrD6d5i)
  - is a compiler that tries to minimise or maximise some attributes of an executable computer program  
  - most common requirement is to minimise the time taken to execute a program, a less common one is to minimise the amount of memory required  
  - it has been shown that some code optimisation problems are NP-complete, or even undecidable

## [LLVM ✨](https://my.mindnode.com/CVQ31TmsUVECWFwFLTRafvy1RbjTzP5E3gjBvs77)

- is compiler infrastructure project which is a ‘collection of modular and reusable compiler and toolchain technologies’ used to develop computer front ends and back ends  
- it is written in C++ and is designed for compile-time, link-time, run-time and ‘idle-time’ optimisations of programs written in arbitrary programming languages  
- it can provide the middle layers of a complete compiler system, taking intermediate representation (IR) code from a compiler and emitting an optimised IR  
	- this new IR can then be converted and linked into machine-dependent assembly language code for a target platform

## [code generation](https://my.mindnode.com/uyVn9qAxzQVsTRdF8cgrJ6ABbYnsJ3LzzTUYxGzG)

- is the process by which a compiler’s code generator converts some intermediate representation of source code into a form (e.g. machine code) that can be readily executed by a machine  
- the input to the code generator typically consists of a parse tree or an abstract syntax tree  
	- the tree is converted into a linear sequence of instructions, usually in an intermediate language such as three-address code

