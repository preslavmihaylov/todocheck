
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


// @fix: malformed todo
// @fix 1: valid todo
// @fix 2: The issue is closed
// @fix 3: The issue is non-existent

# @fix: malformed todo
# @fix 1: valid todo
# @fix 2: The issue is closed
# @fix 3: The issue is non-existent


echo 'this is a simple string with a // @fix: comment in it';

echo 'You can also have embedded newlines in
strings this way as it is
okay to do. But you shouldn\'t match inline // @fix: comments';

/*
 * @fix: Multi-line invalid todo
 */

/*
 * @fix 1: Multi-line valid todo
 */

/*
 * @fix 2: issue is closed
 */

/*
 * @fix 3: issue is non-existent
 */

/**
 * @fix: docstring invalid todo
 */

/**
 * @fix 1: todo is a valid one
 */

/**
 * @fix 2: issue is closed
 */

/**
 * @fix 3: issue is non-existent
 */
