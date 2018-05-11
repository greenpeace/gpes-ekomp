/*
.

Compare two csv files to check which email addresses are in common and which email addresses are missing in each file.

WHAT'S THE PURPOSE OF THIS SCRIPT?

- We needed a quick-to-use database debugging tool for a vey specific problem: to help in a daily check about the information being sent from the mailing program to the CRM, and in the inverse direction.

- This script creates lists with the common and missing email addresses to be investigated. It can be useful for other purposes as well.

- Besides unique email addresses it can compare sha256 hashes or Spanish ID numbers.

- Parsing and comparing a 41MB file with 400.000 email addresses against another 36MB file with 350.000 email addresses takes 9 seconds in a fast laptop.


GET HELP

	./ecompare --help

HOW TO USE IT

Export CSV files

You'll need to export csvs with supporter data in Engaging Networks and create similar exports in Salesforce.

Compare CSV files

Download both files as csvs to the same folder as the script. Then, using the command line run the script as:

    ./ecompare -data=emails -A=fileA.csv -B=fileB.csv

    -data specifies the type of data to compare. It can be emails, sha256, urls or dni (Spanish ID numbers)

    -A and -B specify the names of both files.

Get details about the comparison

When running the script you'll get a quick report like this:

WHAT HAPPENED?

File A: fileA.csv

File B: fileB.csv

Parsed emails in fileA.csv : 229

Parsed emails in fileB.csv: 214

In fileB.csv but not in fileA.csv : 0 emails

In fileA.csv but not in fileB.csv : 15 emails

In both fileA.csv and fileB.csv : 214 emails

And the script always creates 3 files in the current folder with the results:

in-a-but-not-in-b.txt

in-b-but-not-in-a.txt

in-both-a-and-b.txt

The filenames describe it's content and running the script again will overwrite this 3 files.

Open this files with your plain/code text editor and investigate the inconsistencies in both your CRM and mailing programs.

To delete the 3 files created by ecompare:

    ./ecompare -trash

.
*/
package main
