# Program Usage

This program reads a file containing a list of integers, sorts them using a binary search tree (BST), and writes the sorted array to a new file.

## Prerequisites

- Go programming language (Golang) must be installed on your system.
- A text file containing a list of integers. The file should be in a readable format.

## Getting Started

1. Clone the repository or download the program file.
2. Open the terminal or command prompt and navigate to the directory where the program is located.

## Running the Program

1. Open the `main.go` file in a text editor.
2. Update the following lines in the `main` function:

```go
content, err := ioutil.ReadFile("Read File")
```

Replace `"Read File"` with the path to the file you want to read and sort. For example, if the file is located at "C:/Users/Dafi/Documents/VSCode/RandomNumber.txt", the line should be:

```go
content, err := ioutil.ReadFile("C:/Users/Dafi/Documents/VSCode/RandomNumber.txt")
```

```go
err = ioutil.WriteFile("Extracted File", []byte(result), 0644)
```

Replace `"Extracted File"` with the path and name of the file where you want to write the sorted array. For example, if you want to write the sorted array to "C:/Users/Dafi/Documents/VSCode/Sorted.txt", the line should be:

```go
err = ioutil.WriteFile("C:/Users/Dafi/Documents/VSCode/Sorted.txt", []byte(result), 0644)
```

3. Save the changes to the `main.go` file.

4. In the terminal or command prompt, run the following command:

```shell
go run main.go
```

5. The program will read the file, sort the numbers, and write the sorted array to the specified file.

6. After the program finishes, you will see the message:

```
Sorted Array has been written to file. Time elapsed: <elapsed_time>
```

where `<elapsed_time>` is the time taken to sort the array and write it to the file.

## Example

Let's assume we have a file named "RandomNumber.txt" containing the following numbers:

```
10
5
7
3
8
```

We want to sort these numbers and write the sorted array to a file named "Sorted.txt".

To run the program with these inputs, we would make the following changes in the `main` function:

```go
content, err := ioutil.ReadFile("RandomNumber.txt")
```

and

```go
err = ioutil.WriteFile("Sorted.txt", []byte(result), 0644)
```

After running the program, the sorted array will be written to the "Sorted.txt" file, and the output message will be displayed:

```
Sorted Array has been written to file. Time elapsed: <elapsed_time>
```

where `<elapsed_time>` will be the actual time taken by the program to execute.

Please make sure to provide the correct file paths and names according to your system configuration.
