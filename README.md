# ipAddressCounter

## Problem: 
A large file that contains multiple strings (ip addresses) has to be processed in efficient time. We need to count total amount of the unique strings.

## Solution
Simple solution that utilizes a buffer and hash to store addresses. It does not show an improved results, I've considered several different implementations, such as trying to use divide and conquere approach (divide file in multiple files, then use priority queue to count unique addresses across multiple files), using goroutines and other algorithms, but I had a worse result than an algorithm with hash.