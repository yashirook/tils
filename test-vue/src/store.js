import Vue from 'vue';
import axios from 'axios';

const URL_BASE = '/books/'

module.exports = new Vue({
  data: {
    book_list: []
  },
  methods: {
    get_ajax(url, name) {
      return axios.get(URL_BASE + url)
      .then((res) => {
        Vue.set(this, name, res.data);
        this.$emit('GET_AJAX_COMPLETE');
      });
    },
    get_data(name) {
      return this.$data[name];
    }
  }
});