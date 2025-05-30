# Overview

This repository contains a collection of Go-based small programmes which will be helpful for interviews

## Structure

```
├── cancelNotifyWorker
│   └── main.go
├── generate-metamask-signature
│   ├── dist
│   │   ├── bundle.js
│   │   ├── index.html
│   │   └── src
│   │       └── index.html
│   ├── package.json
│   ├── README.md
│   ├── src
│   │   ├── index.html
│   │   └── index.js
│   ├── webpack.config.js
│   └── yarn.lock
├── maxListFrequency
│   └── main.go
├── printNosInSequence
│   └── main.go
├── printPrimeNoSlice
│   └── main.go
├── printTreeStructure
│   └── main.go
└── README.md
```

---

## cancelNotidyWorker

Notify one goroutine when another completes its task using a channel. Once notified, the second goroutine should gracefully stop execution.

**Example Input and Output:**
```
--->  0
printing id -->  0
--->  1
printing id -->  1
...
--->  9
cancelling channel --->  0
Received cancel signal from 0. Exiting...  1
```

[Source Code](./cancelNotidyWorker/main.go)

--- 
## generate-metamask-signature

This project allows users to generate a MetaMask signature in the browser using Webpack and JavaScript. It contains:
- A bundled JavaScript build (`bundle.js`)
- Frontend HTML interface

[Source Code](./generate-metamask-signature/README.md)

---

## maxListFrequency

**Problem Statement:**  
Given an array of integers, find the element that occurs the least number of times. If there are multiple, return the one that appears later in the list.

**Example Input and Output:**
```
Input:  [1, 1, 2, 3, 3, 4]  -> Output: 2
Input:  [1, 2, 2, 3, 3, 4]  -> Output: 4
```

[Source Code](./maxListFrequency/main.go)

---

## printNosInSequence

**Problem Statement:**  
Print numbers in a strict alternating sequence (odd, even) using goroutines and channels.

**Sample Run:**
```
Hello, Print How nos you want to print in Sequnce?: 5
Odd: 1
Even: 2
Odd: 3
Even: 4
Odd: 5
```

[Source Code](./printNosInSequence/main.go)

---

## printPrimeNoSlice

**Problem Statement:**  
Generate and print the first `n` prime numbers using goroutines.

**Sample Run:**
```
Enter the number: 5
2
3
5
7
11
```

[Source Code](./printPrimeNoSlice/main.go)

---

## printTreeStructure

**Problem Statement:**  
Given a list of file paths, print a visual tree structure representing directories and files.

**Example Input:**
```
/app/components/header
/app/services
/app/tests/components/header
/app/components/services
/images/image.png
/tsconfig.json
/index.html
```

**Output:**
```
-app
--components
---header
---services
--services
--tests
---components
----header
-images
--image.png
-tsconfig.json
-index.html
```

[Source Code](./printRreeStructure/main.go)

---

## How to Run
1. Make sure you have Go installed: https://go.dev/dl/
2. Navigate to any sub-folder with a `main.go` file.
3. Run the Go file:

```bash
cd printPrimeNoSlice
go run main.go
```

---

## Contributions
PRs are welcome! Please open an issue first to discuss what you'd like to change.
