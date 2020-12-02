package main

// A malformed TODO comment
// TODO 123: This is a valid todo comment
// TODO 321: This is an invalid todo, marked against a closed issue
/*
 * TODO 567: This is an invalid multiline todo, marked against a non-existent issue
 */

// A malformed ToDo comment
// ToDo 123: This is a valid todo comment
// ToDo 321: This is an invalid todo, marked against a closed issue
/*
 * ToDo 567: This is an invalid multiline todo, marked against a non-existent issue
 */

// A malformed @fixme comment
// @fixme 123: This is a valid todo comment
// @fixme 321: This is an invalid todo, marked against a closed issue
/*
 * @fixme 567: This is an invalid multiline todo, marked against a non-existent issue
 */

// A malformed @fix comment
// @fix 123: This is a valid todo comment
// @fix 321: This is an invalid todo, marked against a closed issue
/*
 * @fix 567: This is an invalid multiline todo, marked against a non-existent issue
 */
