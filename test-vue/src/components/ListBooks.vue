<template>
  <div id="app">
    <ul>
      <li v-for="book in books" :key="book.name">{{book.name}}</li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ListBooks',
  data() {
    return {
      loading: true,
      errored: false,
      error: null,
      todos: null
    };
  },
  method: {
    window:onload = function() {
      axios.get("http://yashiroken.work/books/list")
        .then(response =>{
          this.books = response.data;
        })
        .catch(err => {
          (this.errored = true), (this.error = err);
        })
        .finally(() => (this.loading = false));
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
