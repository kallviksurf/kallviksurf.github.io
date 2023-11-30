#!/bin/bash
 
# Check if `tidy` command-line tool is installed
if ! command -v tidy &> /dev/null
then
    echo "Error: tidy is not installed. Please install tidy using 'brew install tidy' command." >&2
    exit 1
fi
 
# Check if input folder argument is provided
if [ $# -eq 0 ]
then
    echo "Error: No input folder provided. Please provide an input folder as an argument." >&2
    exit 1
fi
 
# Check if input folder exists
if [ ! -d "$1" ]
then
    echo "Error: Folder '$1' does not exist." >&2
    exit 1
fi
 
# Loop through HTML files in the input folder and beautify them
find "$1" -type f -name "*.html" -print0 | while read -d $'\0' file
do
    # Beautify HTML code and replace the original file
    tidy -indent -wrap 0 -quiet -m -i "$file"
done
 
# Exit with success status
exit 0
