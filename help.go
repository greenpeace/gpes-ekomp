package main

import "fmt"

func helpMe() {

	textToPrint := `

* * * HELP * * *

Compares specific data in a text file with an encrypted list.

- specific data includes emails, or Spanish ID numbers (DNI)
- text files can be .csv, .txt, .sql or .html

Use as in the 2 following examples:

./ekomp -data=emails -file=myfile.txt -list=encrypted.txt
./ekomp -data=dni -file=myfile.txt -list=encrypted.txt

This script always creates 2 files in the current folder with the results: 

1 - was-found.txt
2 - was-not-found.txt

Each time the script runs it overwrites this 2 files.

---------------------

Comand line options:

-help			Display this help
-data=emails		What to compare in the files. It can be "emails" or "dni". By default it compares emails.
-file=myfile.txt	File to check
-list=encrypted.txt	File with the encrypted data
-debug=true		Debug the script					
-trash			To delete the 2 files created by ekomp

`
	fmt.Printf(textToPrint)

}
