#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <directory_path>"
    exit 1
fi

directory_path="$1"

if [ ! -d "$directory_path" ]; then
    echo "Error: The path provided is not a directory."
    exit 2
fi


cd "/home/cyrus/trivil/src" \
&&
./trivil -out "$directory_path/generated" "$directory_path" \
&& \
cd "$directory_path/generated" \
&& \
/usr/bin/java -jar /home/cyrus/Downloads/jasmin-2.4/jasmin.jar *.j \
&& \
echo -e "\n==================== OUTPUT ======================" \
&& \
/usr/bin/java Main * \
&& \
echo -e "\n==================================================="
