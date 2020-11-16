

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


# This is a single-line malformed @fix

"""
And this is a multiline malformed @fix
It should be parsed properly
"""

'''
This is the same multiline malformed @fix
but with single-quotes
'''

myvar = 5 # This is a malformed @fix at the end of a line

# @fix 1: This is a valid todo comment

hello = "hello" # @fix 234: This is an invalid todo, with a closed issue

"""
@fix 234: This is an invalid todo, marked against a closed issue
"""

'''
@fix 234: This is an invalid todo,
marked against a closed issue with single quotes
'''
