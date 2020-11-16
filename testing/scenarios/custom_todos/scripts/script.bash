# This is in a bash script

# A malformed TODO comment
# TODO 123: This is a valid todo comment
# TODO 321: This is an invalid todo, marked against a closed issue
curl "localhost:8080" # TODO 567: This is an invalid todo, marked against a non-existent issue

# A malformed ToDo comment
# ToDo 123: This is a valid todo comment
# ToDo 321: This is an invalid todo, marked against a closed issue
curl "localhost:8080" # ToDo 567: This is an invalid todo, marked against a non-existent issue

# A malformed @fixme comment
# @fixme 123: This is a valid todo comment
# @fixme 321: This is an invalid todo, marked against a closed issue
curl "localhost:8080" # @fixme 567: This is an invalid todo, marked against a non-existent issue

# A malformed @fix comment
# @fix 123: This is a valid todo comment
# @fix 321: This is an invalid todo, marked against a closed issue
curl "localhost:8080" # @fix 567: This is an invalid todo, marked against a non-existent issue
