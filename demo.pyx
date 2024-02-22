#!python
#cython: language_level=3

from libc.stdio cimport printf

# a cython C function that calls some python code and requires python interpreter to be initialized
cdef public int hello():
     printf("printed using cython's libc.stdio.printf\n")
     print("printed using python's print")
     return 42

