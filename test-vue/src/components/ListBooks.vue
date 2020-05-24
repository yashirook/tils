<template>
  <div id="booklist">
    <ul v-for="book in books" :key="book.name">
      <li>
        書名:{{ book.name }}, コメント:{{ book.comment }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "booklist",
  data() {
    return {
      books: null
    };
  },
  created() {
    axios
      .get("/books/list")
      .then(response => {
        this.books = response.data;
      })
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
