//TODO: regular inline comment

/*
* block comment
*/

/**
* groovydoc comment
*/

/*
* TODO: Multi-line invalid todo
*/

/**
* TODO: groovydoc invalid todo
*/

// TODO 2: The issue is closed

// TODO 3: The issue is non-existent

def trippleDoubleQuotedString = """string start
    //TODO: inline comment inside a string block. Shouldn't be captured.

    /* block comment inside a string block */
    /** groovydoc comment inside a string block */
string end"""

def trippleSingleQuotedString = '''string2 start
    //TODO: inline comment inside a string2 block. Shouldn't be captured

    /* block comment inside a string2 block */
    /** groovydoc comment inside a string2 block */
string2 end'''

def singleQuotedString = 'string2 start // TODO: inline string shouldn\'t be captured'
def doubleQuotedString = "string2 start // TODO: inline string \"shouldn't\" be captured"

println trippleDoubleQuotedString
println ''
println trippleSingleQuotedString

/**
* TODO 1: groovydoc valid todo
*/

/*
* TODO 1: multi-line valid todo
*/

/*
* TODO 2: Invalid todo as issue is closed
*/

/**
* TODO 2: Invalid todo as issue is closed
*/


//@fix: regular inline comment

/*
* block comment
*/

/**
* groovydoc comment
*/

/*
* @fix: Multi-line invalid todo
*/

/**
* @fix: groovydoc invalid todo
*/

// @fix 2: The issue is closed

// @fix 3: The issue is non-existent

def trippleDoubleQuotedString = """string start
    //@fix: inline comment inside a string block. Shouldn't be captured.

    /* block comment inside a string block */
    /** groovydoc comment inside a string block */
string end"""

def trippleSingleQuotedString = '''string2 start
    //@fix: inline comment inside a string2 block. Shouldn't be captured

    /* block comment inside a string2 block */
    /** groovydoc comment inside a string2 block */
string2 end'''

def singleQuotedString = 'string2 start // @fix: inline string shouldn\'t be captured'
def doubleQuotedString = "string2 start // @fix: inline string \"shouldn't\" be captured"

println trippleDoubleQuotedString
println ''
println trippleSingleQuotedString

/**
* @fix 1: groovydoc valid todo
*/

/*
* @fix 1: multi-line valid todo
*/

/*
* @fix 2: Invalid todo as issue is closed
*/

/**
* @fix 2: Invalid todo as issue is closed
*/
