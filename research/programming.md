# version control


## [mercurial](https://my.mindnode.com/NxcTDxweypHJ5dkqR7bWpryPxjbmzpCAxcKV9tG4)

## [git ✨](https://my.mindnode.com/TN2aYRpMdhpZfcqqCeaK4tbSjaVYFDJupN9UYZBS#1008.9,-448.6,0)


# [open source ✨](https://my.mindnode.com/fUS8UzoVqq438SCBhimjPJxwKs1YEw2Te2mPhDdi)

what it means to work with open source and how to best sustain an open source project


# practice


## [programming](https://my.mindnode.com/XVKPw7bLpLTzXUnqPf8QYU7mNVxwxdq8Zc316wJQ)

### [katas](https://my.mindnode.com/duoYeBBx9GzeasYxzASYjzgzvqfqQNsw37ZzaJ5V)


# [programming ✨](https://my.mindnode.com/nyxpNazrso6eyV3M4RNGS9sUyXfNzTXq6pGhRYLZ#-1120.9,46.4,-1)

- is a process that leads from an original formulation of a computing problem to an executable computer programs


## [bitwise operations](https://my.mindnode.com/jjdFxBtm7BCMmZBfFHwwyRh3q3awMFxbXePMs4Cy)

- operate on one or more bit patterns or binary numerals at the level of individual bits  
- it is a fast, simple action directly supported by the processor, and is used to manipulate values for comparisons and calculations  
- the operators are :  
	- AND  
	- OR  
	- NOT  
	- XOR (Exclusive Or)  
	- NAND (Not And)

## [design patterns](https://my.mindnode.com/FPsEgZQpPDCmzQSspyY2AKqGNTsLGrAKpg6sXk12#1041.1,-95.8,0)

- are general reusable solutions to a commonly occurring problem within a given context in software design  
- it is not a finished design that can be transformed directly into source or machine code  
	- it is a description or template for how to solve a problem that can be used in many different situations

## [dynamic programming](https://my.mindnode.com/cpxi9EqMbbF7cvq3vqD4pDQLwHetX81qagPEXi99)

- is a method for solving a complex problem by breaking it down into a collection of simpler subproblems, solving each of those subproblems just once, and storing their solutions - ideally, using a memory-based data structure  
	- next time the same subproblem occurs, instead of recomputing its solution, one simply looks up the previously computed solution, thereby saving computation time at the expense of a (hopefully) modest expenditure in storage space  
- the technique of storing solutions to subproblems instead of recomputing them is called ‘memoization’  
- dynamic problem algorithms are often used for optimisation  
  
method of solving problems, using this method, a complex problem is split into simpler problems, which are then solved, at the end, the solutions of the simpler problems are used to find the solution of the original complex problem

## [paradigms](https://my.mindnode.com/qqjRxmiDoMyKteYicvarwjrgMgpsYJjLWgXsD5Wn#767.7,-730.8,2)

- are a way to classify programming languages according to the style of computer programming   
- features of various programming languages determine which programming paradigms they belong to  
	- as a result, some languages fall into only one paradigm, while others fall into multiple paradigms  
- common paradigms include  
	- imperative which allows side effects  
	- functional which disallows side effects  
	- declarative which does not state the order in which the operations execute  
	- object-oriented which groups code together with the state the code modifies  
	- procedural which groups code into functions  
	- logic which has a particular style of execution model coupled to a particular style of syntax and grammar  
	- symbolic programming which has a particular style of syntax and grammar

### declarative

- is a style of building the structure and elements of computer programs that expresses the logic of a computation without describing the control flow

- [functional ✨](https://my.mindnode.com/yLFW2vMCgshzaepsYcYDCpYnRBFktycVYmjtRwcF)
  - is a style of building the structure and elements of computer programs that treats computation as the evaluation of mathematical functions and avoids changing state and mutable data  
  - in functional code, the output value of a function depends only on the argument that are passed to the function, so calling a function f twice with the same value for an argument x will produce the same result f(x) each time  
  - eliminating side effects (changes in state that do not depend on the function inputs), can make it much easier to understand and predict the behaviour of a program  
  - lambda calculus provides a theoretical framework for describing functions and their evaluation

- logic
  - is largely based on formal logic  
  - any program written in a logic programming language is a set of sentences in logical form, expressing facts and rules about some problem domain

- dataflow
  - models a program as a directed graph of the data flowing between operations, thus implementing data flow principles and architecture

- constraint
  - is a paradigm wherein relations between variables are stated in the form of constraints  
  - constraints differ from the common primitives of imperative programming languages in that they do not specify a step or sequence of steps to execute, but rather the properties of a solution to be found  
  - constrain programming can be expresses in the form of constrain logic programming, which embeds constrains into a logic program

### imperative

- uses statements that change a program’s state  
- imperative program consists of commands for the computer to perform  
	- it focuses on describing how a program operates

- [object oriented](https://my.mindnode.com/bXt6DgouTpBrCURxSzAyvZxMEsSU4pd4ecBVtRzn)
  - is based on the concepts of ‘objects’, which may contain data, in the form of fields, often known as attributes; and code, in the form of procedures, often known as methods  
  - a feature of objects is that an object’s procedures can access and often modify the data fields of the object with which they are associated (objects have a notion of ‘this’ or ‘self’)  
  - in OOP, computer programs are designed by making them out of objects that interact with one another  
  - most popular object oriented languages are class-based, meaning that objects are instances of classes, which typically also determine their type

- procedural
  - is derived from structured programming, based upon the concept of the procedure call  
  - procedures, also known as routines, subroutines, or functions, simply contain a series of computation steps to be carried out  
  - any given procedure might be called at any point during a program’s execution, including by other procedures or itself

### metaprogramming

- is a programming technique in which computer programs have the ability to treat programs as their data  
- it means that a program can be designed to read, generate, analyse or transform other programs, and even modify itself while running  
- it can be used to move computations from run-time to compile-time, to generate code using compile time computations, and to enable self-modifying code

- [template metaprogramming](https://my.mindnode.com/GkpWKLpuRFrkF6zPEW92CzDtyaxua4s7YyAmZ4Kn)
  - is a metaprogramming technique in which templates are used by a compiler to generate temporary source code,, which is merged by the compiler to generate temporary source code, which is merged by the compiler with the rest of the source code and then compiled  
  - the output of these templates include compile-time constants, data structures and complete functions  
  	- the use of templates can be thought of as compile-time execution

## [testing](https://my.mindnode.com/1qGKADJWJ2MtkLpVgCUF5Pooa4iVKbHjTfKfYGNu)

- software testing is an investigation conducted to provide stakeholders with information about the quality of the product or service under test  
- it involves the execution of a software component or system component to evaluate one or more properties of interest  
	- it indicates the extent to which the component or system under test  
		- meets the requirements that guided its design and development  
		- responds correctly to all kinds of inputs  
		- performs its functions within an acceptable time  
		- is sufficiently usable  
		- can be installed and run in its intended environments  
		- achieves the general result its stakeholders desire

### [test driven development](https://my.mindnode.com/ey6QYxyUo6JqQ6tDmmRRLykwjpr7ybbEf5fcicZr)

- is a software development process that relies on the repetition of a very short development cycle  
	- requirements are turned into very specific test cases, then the software is improved to pass the new tests, only  
- it encourages simple designs and inspires confidence   
- a typical test driven cycle is  
	1. add a test  
	2. run all tests and see if the new test fails  
	3. write the code  
	4. run tests  
	5. refactor code

### [unit testing](https://my.mindnode.com/z3wLxo6zw1KEpPs1kryj51ToaDH3RkfEDQx4GRDR)

- is a software testing method where units of source code, sets of one or more computer program modules together with associated control data, usage procedures, and operating procedures, are tested to determine whether they are fit for use  
- one can view a unit as the smallest testable part of an application  
- ideally, each test case is independent from the others

## [exception handling ✨](https://my.mindnode.com/xEiGQnjcJeqyxvxpxspGJR7q4zDQPawRYMpJDqVj)

- exception handing is the process of responding to the occurrence, during computation, of exceptions (anomalous or exceptional conditions requiring special processing - often changing the normal flow of program execution)  
- in general, an exception breaks the normal flow of execution and executes a pre-registered exception handler

## [abstraction](https://my.mindnode.com/7aNxJDAnKfSqGHhHpnpewSeDdFMmCY5tj32fdFqH)

- is a technique for arranging complexity of computer systems  
- it works by establishing a level of complexity on which a person interacts with the system, surpassing the more complex details below the current level  
- the programmer works with an idealised interface (usually well defined) and can add additional levels of functionality that would otherwise be too complex to handle  
- abstraction can apply to control or to data  
	- control abstraction is the abstraction of actions while data abstractions is that of data structures  
	- control abstraction involves the use of subroutines and control flow abstractions  
	- data abstraction allows handling pieces of data in meaningful ways  
- the notion of an object in object-oriented programming can be viewed as a way to combine abstractions of data and code

## [software verification](https://my.mindnode.com/i2Tmz28SzbT3WCq6aqrQy75sy226sNxAkZ3aABxG)

- is a discipline of software engineering whose goal is to issue that software fully satisfies all the expected requirements  
- there are two fundamental approaches to verification   
	- dynamic verification, also known as Test or Experimentation (this is good for finding bugs)  
	- static verification, also known as Analysis (is useful for proving correctness of a program although it may result in false positives)

## [dependency injection](https://my.mindnode.com/L1ZfeUJUTiLaJG9RUbE47q6QkDkYUxBGMA9URp2E)

- is a technique whereby one object supplies the dependencies of another object  
- a dependency is an object that can be used (a service)  
- an injection is the passing of a dependency to a dependent object (a client) that would use it  
	- the service is made part of the client’s state  
	- passing the service to the client, rather than allowing a client to build or find the service, is the fundamental requirement of the pattern  
- the intent behind it, is to decouple objects to the extent that no client code has to be changed simply because an object it depends on needs to be changed to a different one

## [syntax trees](https://my.mindnode.com/55ZLtTbzDZJV6aXqXUxN92g6Ps7GMbnsEpjD7sqq)

- are tree representations of the abstract syntactic structure of source code written in a programming language  
- each node of the tree denotes a construct occurring in the source code  
- they are used in program analysis and program transformation systems  
- they are data structures widely used in compilers due to their property of representing the structure of program code  
	- an AST is usually the result of the syntax analysis phase of a compiler  
	- it often serves as an intermediate representation of the program through several stages that the compiler requires, and has a strong impact on the final output of the compiler

## [static analysis](https://my.mindnode.com/3sCkLyCxbePrDHzrCRpVEQX6xfMHFPSrN6tRzMfx)

- is the analysis of computer software that is performed without actually executing programs (analysis performed on executing programs is known as dynamic analysis)  
	- in most cases the analysis is performed on some version of the source code, and in the other cases, some form of the object code

## [linters](https://my.mindnode.com/6KxuVfowmW81Teuawc2FedkqG5QxiHwseUg3GZDP)

- linting is the process of running a   
program that will analyse code for potential errors

## [evaluation](https://my.mindnode.com/NDkKMcxWGi5vHMwEbFzpM3yof8orjxvCyPuGUvCn)

- is a systematic method for collecting, analysing and using information to answer questions about projects, policies and programs, particularly about their effectiveness and efficiency 

## [memory](https://my.mindnode.com/wbRrdVuWerWQMHiPNrW7KdeLc9WFyz2EzBBQSgSd)

- how to effectively manage data whilst programming

## [type system ✨](https://my.mindnode.com/S5yqn46mnWuPPzHzQN5yaZPzoJ2gCfTbJWzJt4qZ)

- is a set of rules that assign a property called a type to various contracts a computer program consists of, such as variables, expressions, functions or modules  
- the main purpose of it is to reduce possibilities for bugs in computer programs by defining interfaces between different parts of a computer program, and then checking that the parts have been connected in a consistent way  
	- this checking can happen   
		- statically (at compile time)  
		- dynamically (at run time)  
		- or as a combination of static and dynamic checking  
- other purposes of a type system include enabling certain compiler optimisations, allowing for multiple dispatch, providing a form of documentation, etc.  
- it associates a type with each computed value and, by examining the flow of these values, attempts to ensure or prove that no type errors can occur

## [competitive ✨](https://my.mindnode.com/HjzAcCjnzpEcL69NSNYigoEyVeDTUiXLsEh3nM6P)

- is a mind sport usually held over internet or a local network which involves participants trying to program according to provided specifications  
- the competition generally involves a set of logical or mathematical problems and contestants are required to write computer programs capable of solving each problem

## debugging

## [language design](https://my.mindnode.com/SFyx1AMA4eKBbkwq59ymrQz3AEKQMiw8A8SV1WeX)


# [emulation](https://my.mindnode.com/wKPccGXqDxbUMndMp1nPVs1NbzhzcmsqKzBDyXWx#477.8,-480.0,2)


# task runners


## [gulp](https://my.mindnode.com/btgJru7wziUq7q2jYLaTe3cpe9CSsXLZpAPJ2MF8)

## grunt


# [virtual machines](https://my.mindnode.com/yYb2h4qJrcC6xq553ubAqbW1hg7dTXV7Xxpz13vn)

a program on a computer that works like it is a separate computer inside the main computer, the program that controls visual machines is called a hypervisor and the computer that is running the virtual machine is called the host


## [jvm ✨](https://my.mindnode.com/xJivER7pcpjTAHsCwjtpd9meWnzLeMLLYkGLzQZk)

is a set of computer software programs and data structures which use a virtual machine model for the execution of other computer programs and scripts, operates on Java bytecode

### [vert.x](https://my.mindnode.com/m93VE11WC2DEZzbH2ztmdqow9yZGom3VBNGoqFAv)

## [clr](https://my.mindnode.com/TzP7AgGyszyGkMpSKtz1YqPCpytLxFQj2JpMSqsk)

## [hypervisor](https://my.mindnode.com/iEH6suLAhq52zp4GjLNFLbAgm9YQ5BKkaUY3Cx8T)


# system design

- is the process of defining the architecture, components, modules, interfaces and data for a system to satisfy specified requirements


## [software architecture ✨](https://my.mindnode.com/dCGPsRx5okxjPJByJqDC6Kfbz3J9PpqeR55xmRvj)

- referee to the fundamental structures of a software system, the discipline of creating such structures, and the documentation of these structures  
- it is about making fundamental structural choices which are costly to change once implemented


# [software development](https://my.mindnode.com/WgXAsAKvd4MsMLqxwqvbHuMVz4YYXqVkvZVtSwXB)


## [agile](https://my.mindnode.com/4hPUYtVtfLEL17EEecXhszyRpiM1pmjYec3i9ytr)

describe a set of principles for software development under which requirements and solutions evolve through the collaborative effort or self organising cross functional teams, it advocates adaptive planning, evolutionary development, early delivery and continuous improvement and it encourages rapid and flexible response to change

## [devops ✨](https://my.mindnode.com/ns3bCv8vPJrsDaq5koMfMsBZzJfRgeYqmy7SusxR)

refers to a set of practices that emphasise the collaboration and communication of both software developers and information technology (IT) professionals while automating the process of software delivery and infrastructure change, aims to establishing a culture and environment where building, testing and releasing software can happen rapidly, frequently and more reliably 

## approaches

### [domain driven design ✨](https://my.mindnode.com/bCGsYrxxdgxfSS6x8WCHJ8FMMupN9LuHNKcybeKN)

approach to software development for complex needs by connecting the implementation to an evolving model, it follows :   
  
1. placing the project’s primary focus on the core domain and logic  
2. basing complex designs on a model of the domain  
3. initiating a creative collaboration between technical and domain experts to iteratively refine a conceptual model that addresses particular domain problems


# [system administration](https://my.mindnode.com/YtbeRt3uZR2o3b8R3fx58i6YomayqJEKwP5z8JE5)


# [cellular automata](https://my.mindnode.com/7WxM6UNViGoDE4qpgePBpQyyWURAuSzVLa7zQzEL)

- models a dynamic system by using a number of cells where each cell has one of several possible states, with each “turn” or iteration the state of the current cell is determined by two things   
	1. its current state  
	2. states of the neighbouring cells


## [game of life](https://my.mindnode.com/sBQxfKxmxqg1LqSCN5pBFkjqJvcQuFjyifqTNHtk)

- the ‘game’ is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input  
- rules of the game are : 	  
	1. any live cell with fewer than two live neighbours dies, as if caused by underpopulation  
	2. any live cell with two or three live neighbours lives on to the next generation  
	3. any live cell with more than three live neighbours dies, as if by overpopulation  
	4. any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction


# [configuration management](https://my.mindnode.com/gUX5gmQUmdzo9MWcyggzYBzRF75PpE16ob1xjyMs)

is a systems engineering process for establishing and maintaining consistency of a product’s performance, functional and physical attributes with its requirements, design and operational information through its life


## [puppet](https://my.mindnode.com/8HRdnRDFCxspXpf7qdGpe84TipqnasfkHQYCcf5y)

## [ansible](https://my.mindnode.com/rWEgABaoDTqViLJ6Payck4p4S8sb8jx2F1H81CAa)

## [chef](https://my.mindnode.com/AqwnbN1j14mSn27aqwz3EDnpmFdX5essxpJKpU7H)


# build systems


## [cmake ✨](https://my.mindnode.com/nrk4ex9w31toZU8dvztZq76ipyfs7ymsDXDzJFwg)


# [serverless architecture ✨](https://my.mindnode.com/hqjmqW2xMYTekHSzwDpApeqqfsyX2M1SAakkHKSN)


# [software deployment ✨](https://my.mindnode.com/JHNueMJWFJbZJiHzQdzuY5G8TWBZW4xzEKU19xTx)


## travis CI

## [microservices](https://my.mindnode.com/u79EngB79DqdaYsAzDkGycjBWGAphxDbZp1JpyEx)

## [kubernetes ✨](https://my.mindnode.com/YJ9t85qWRsBsbnJCxj3rvTZPSsrXs62qmKbLv97s)

## [docker ✨](https://my.mindnode.com/3tDDSeJdpXeWHxU2L4mLyzpeAGLmCxepmMLWyu1R)


# [distributed computing](https://my.mindnode.com/sEKHQbwqA8NCpJWSSttobDKJrLvoppdBtafUH17X)


## [petri nets](https://my.mindnode.com/szji6NR1NPFanHbqpopkPqgTiPc3gAzehqpNX5NL)

## [hadoop](https://my.mindnode.com/nFdNGaPw1BgCsyiDnMPtSwtC8ZBJqHQAEVsnghXz)

## [apache spark](https://my.mindnode.com/VNVjTVxwpDqC2rdvLwPk8HS5GBqf2rEQMxW1Bi3q)

