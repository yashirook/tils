<template>
    <div class="registerbook">
        <h2>Bookshelf</h2>
        <input type="text" placeholder="Bookname" v-model="bookname">
        <input type="text" placeholder="Comment" v-model="comment">
        <button @click="registerBook">Register</button>
        <p>Do you have an account?
            <router-link to="/signin">sign in now!!</router-link>
        </p>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RegisterBook',
  data () {
    return {
      bookname: '',
      comment: ''
    }
  },
  methods: {
    registerBook: function () {
      axios.post('/books/list', {
      name: this.bookname,
      comment: this.comment
      })
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
    }
  }
}
</script>

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
.signup {
    margin-top: 20px;
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
}
input {
    margin: 10px 0;
    padding: 10px;
}
button {
    margin: 10px 0;
    padding: 10px;
}
</style>
