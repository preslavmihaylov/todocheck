# This is a single-line malformed TODO

# TODO 1: This is a valid todo comment

when isMainModule:
# TODO 234: Invalid todo, with a closed issue
    discard "Hello World" 

#[ TODO 2: Another valid todo ]#

    discard "Magic here, magic there!"
#[
#[ TODO 3: There is also a valid nested todo here! ]#
]#