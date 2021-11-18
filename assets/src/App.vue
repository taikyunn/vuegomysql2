<template>
  <div id="app">
    <h3>掲示板に投稿する</h3>
    <label for="name">ニックネーム</label>
    <input
      id="name"
      type="text"
      v-model="name"
    >
    <br><br>
    <label for="comment">コメント</label>
    <textarea
      id="comment"
      v-model="comment"
    ></textarea>
    <br>
    <button @click="createComment">コメントをサーバーに送る</button>

    <h2>掲示板</h2>
  </div>
</template>

<script>
import axios from 'axios'

export default {
    data() {
    return {
      name: "",
      comment: "",
    }
  },
  methods: {
    createComment() {
      const params = new URLSearchParams();
      params.append('name', this.name)
      params.append('comment', this.comment)
      axios.post('/new', params)
      .then(response => {
        console.log(response)
      })
      .catch(error => {
        console.log(error)
      })
      this.name = ''
      this.comment = ''
    }
  }
}
</script>

<style>
</style>
