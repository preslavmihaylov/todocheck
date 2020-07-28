
// TODO: malformed todo
// TODO 1: valid todo
// TODO 2: The issue is closed
// TODO 3: The issue is non-existent

# TODO: malformed todo
# TODO 1: valid todo
# TODO 2: The issue is closed
# TODO 3: The issue is non-existent


echo 'this is a simple string with a // TODO: comment in it';

echo 'You can also have embedded newlines in
strings this way as it is
okay to do. But you shouldn\'t match inline // TODO: comments';

/*
 * TODO: Multi-line invalid todo
 */

/*
 * TODO 1: Multi-line valid todo
 */

/*
 * TODO 2: issue is closed
 */

/*
 * TODO 3: issue is non-existent
 */

/**
 * TODO: docstring invalid todo
 */

/**
 * TODO 1: todo is a valid one
 */

/**
 * TODO 2: issue is closed
 */

/**
 * TODO 3: issue is non-existent
 */
