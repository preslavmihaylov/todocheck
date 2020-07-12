

# This is a single-line malformed TODO

"""
And this is a multiline malformed TODO
It should be parsed properly
"""

'''
This is the same multiline malformed TODO
but with single-quotes
'''

myvar = 5 # This is a malformed TODO at the end of a line

# TODO 1: This is a valid todo comment

hello = "hello" # TODO 234: This is an invalid todo, with a closed issue

"""
TODO 234: This is an invalid todo, marked against a closed issue
"""

'''
TODO 234: This is an invalid todo,
marked against a closed issue with single quotes
'''
