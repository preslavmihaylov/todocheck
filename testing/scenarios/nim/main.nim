# This is a single-line malformed TODO

# TODO 1: This is a valid todo comment

when isMainModule:
    # TODO 234: Invalid todo, with a closed issue
    discard "Hello World" # This is a malformed TODO at the end of a line

    #[ TODO 2: Another valid todo ]#

    discard "Magic here, magic there!"
    #[ There is an explaination here, but furthermore
    #[ TODO 3: There is also a valid nested todo here! ]#
    ]#