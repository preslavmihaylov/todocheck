// oneline comment, malformed TODO
// TODO 1: oneline comment wellformed
// TODO 2: oneline comment with Issue closed

<!-- TODO 3: wellformed HTML multiline entry -->
<!-- TODO 2: wellformed HTML multiline entry, BUT issue closed -->
<!-- TODO: missing number -->

/* TODO 1: wellformed CS/JS multiline entry */
/* TODO 2: wellformed CS/JS multiline entry, BUT issue closed */
/*
wellformed CS/JS multline entry
*/

/*
TODO: malformed CS/JS multline entry, missing number
*/

<!--
TODO 4: wellformed HTML multiline entry
-->


"this is a // TODO 1: valid comment in a string"
"this is a /* TODO 1: valid multiline comment in a string */"
"this is a <!-- TODO 1: valid HTML multiline comment in a string -->"

<template>
  <div id="app">
    {{ message }}
  </div>
</template> /* TODO 1: valid multiline comment in code */

<script>
export default {
  data() {
    return {
      message: 'Hello World',
    };
  },
};
</script> <!-- TODO 1: 
valid HTML multiline comment in code -->

<style>
#app {
  font-size: 18px;
  font-family: 'Roboto', sans-serif;
  color: blue; // TODO 2: online comment wellformed in code, BUT issue closed
}
</style>