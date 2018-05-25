# eKomp

**Compares emails or ID numbers in a text file with an encrypted list.**

## What's the purpose?

With **ekomp** you can check if the emails or Spanish ID numbers you have in a file are in an encrypted blacklist. It's like [ecompare](https://github.com/greenpeace/gpes-ecompare) but the list is encrypted for better privacy and security.

Ekomp has better performance than other tools. It takes about 18 seconds to compare a file with 1,600,000 emails against a list of 162,000 encrypted emails. (Using a 4 core laptop)

## How to use ekomp

### Get help

**ekomp** is a command line script. To get help type in the command line:

```bash
./ekomp --help
```

### Create encrypted lists

You should use **[ecounter](https://github.com/greenpeace/gpes-ecounter)** to create an encrypted lists with emails or dnis. 

* Encrypted email lists use sha256 over lowercased email addresses *(your.name@domain.com)*
* Encrypted "DNI" lists use sha256 over uppercased Spanish ID numbers. *(64580957Q or X9137239Y)*

Please note: If you use another tool to create sha256 files don't forget that using uppercase for ID numbers and lowercase for emails is crucial.

### Compare your data with an encrypted list

Put the text file and the encrypted list in the same folder as the script. Then, using the command line run the script as this example:

```bash
./ekomp -data=emails -file=your-file.txt -list=your-sha256-list.txt
```

**Options:**

* `-data` specifies the type of data to compare. It can be `emails` or `dni` (Spanish ID numbers)
* `-file` is the file to check. CSV or txt.
* `-list` is the name of the sha256-encrypted file with the data. This file can be created with, [ecounter](https://github.com/greenpeace/gpes-ecounter).

### Report

When running the script you'll get a quick report like this:

```bash
WHAT HAPPENED?

Your file: your-file.txt
Encrypted file: your-sha256-list.txt
Parsed emails in your-file.txt : 5
Parsed sha256 in your-sha256-list.txt : 4
Was not found in your-file.txt when comparing to your-sha256-list.txt : 1 emails
Was found in your-file.txt when comparing to your-sha256-list.txt : 4 emails
```

### Results

**ekomp** always creates 2 files in the current folder with the results:

1. *was-found.txt*
1. *was-not-found.txt*

The filenames describe it's content and running the script again will overwrite this 2 files.

#### Delete files created by ekomp

To delete the 2 files created by ekomp type:

```bash
./ecompare -trash
```

## Download and use

1. Download the [latest version of the binary code](https://github.com/greenpeace/gpes-ekomp/releases/) for your operating system to your desktop folder.
1. Unzip it to the desktop folder. *(Optionally copy the executable file to a folder in your [path](https://goo.gl/oLzTGw))*
1. To test your install, open the command line, go to the desktop folder and test it with the command:

* `./ekomp --help` *(Mac or Linux)*
* `./ekomp.exe --help` *(Windows)*

## Install from the source code

This script is also provided as [source code](https://github.com/greenpeace/gpes-ecompare/) in [Go](https://golang.org/dl/). To install:

```bash
go get github.com/greenpeace/gpes-ekomp
go install github.com/greenpeace/gpes-ekomp
```


## Note

This script works by parsing text-based files with a regular expression rule to find emails, sha256 hashes or ID numbers. 

The advantage of using regular expressions to parse the files is speed, as the user doesn't has to adjust the files format. 

The disadvantage is that sometimes certain email addresses and DNIs aren't grabbed. Mostly it will be invalid data, but the script report will not match the CRM report.
