# GO-Beginners-Tutorial

############################################################################
In the folder chapters, you can find the basic syntaxes, data types and logic handling for GO
1-basic
How to show output in go

2-variables
How to declare variables in different ways in go

3-arrays
What is an array and slice in arrays in go

4-loops
How to use for and for each loop in go

5-multipath-decision-statement
How to incorporate simple logic and simple validation in go

6-function
How to encapsulate different activities into functions in go

7-maps
How to use the datatype maps in go

8-struct
How to use the user defined datatype struct in go

############################################################################

############################################################################
In chapters/packages you will find how to separate common functions for different go files.
packages. helper.ValidateUserInput is called in main.go from chapters/packages/helper/helper.go. Check the file 
chapters/packages/helper/helper.go for the function that was abstracted and used as a paackage function.

############################################################################

############################################################################

In main.go, concurrency is handled in 
a) loops
b) without loops, but using waiting groups

############################################################################

For first time use:
Create a mode file using the following command where booking-app is the name of the folder module: go mod init booking-app

To run a go file, use the following command where main.go is the is the file to run: go run main.go

To run all the files in a folder, use the following command: go run .
############################################################################
Special thanks to: TechWorld with Nana
