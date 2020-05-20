<template>
  <div id="app">
    <ul>
      <list-item v-for="(value, key) in book_list" :key="name"></list-item>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';
import ListItem from './ListBooks.vue';

export default {
  data() {
    return {
      book_list: []
    };
  },
  components: {
    ListItem
  },
  async created() {
    try {
      let res = await axios.get(URL_BASE + url)
      this.book_list = res.data
    } catch (e) {
      console.error(e)
    }
  }
};
</script>