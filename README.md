# DSA_SEARCHS.

Algorithms for String and Array Searches

This repository contains implementations of various search algorithms in Go programming language.
Algorithms Implemented

    Binary Search
        File: binarySearch/binarySearch.go
        Test: binarySearch/binarySearch_test.go
        Description: Binary search is a classic search algorithm that operates on sorted arrays.

    Linear Search
        File: linearsearch/linearsearch.go
        Test: linearsearch/linearsearch_test.go
        Description: Linear search is a simple search algorithm that sequentially checks each element of the array.

    Interpolation Search
        File: interpolationsearch/interpolation_search.go
        Test: interpolationsearch/interpolation_search_test.go
        Description: Interpolation search improves upon binary search by calculating the position of the target more efficiently for uniformly distributed data.

    Ternary Search
        File: ternarySearch/ternary_serach.go
        Test: ternarySearch/ternary_search_test.go
        Description: Ternary search is a divide and conquer algorithm that works on a sorted array, dividing it into three parts and discarding one part based on the comparison with the key.

    Substring Search
        File: find_substring/find_substring.go
        Test: find_substring/find_substring_test.go
        Description: Substring search checks for the presence of a substring within a given string using simple iteration..

Usage

Each search algorithm is implemented in its own directory along with its corresponding test file. To run the tests for any algorithm, navigate to its directory and execute go test.
Example

To run tests for binary search:



      cd binarySearch
      go test

Contributing

Contributions and improvements are welcome! If you have any suggestions or find issues, please feel free to open an issue or submit a pull request.
