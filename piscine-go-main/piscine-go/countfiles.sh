#!/bin/bash

# Count regular files
file_count=$(find . -type f | wc -l)

# Count directories (including current directory)
dir_count=$(find . -type d | wc -l)

# Sum both counts and print only the total number
echo $((file_count + dir_count))

