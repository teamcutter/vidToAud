<template>
    <div class="main-page">
      <div class="search-form">
        <form @submit.prevent="search">
        <input v-model="query" type="text" placeholder="YouTube link">
        <button type="submit">Search</button>
      </form>
      <div v-if="loading">
        <h1>Waiting...</h1>
      </div>
      <div v-if="audioUrl">
      <a :href="audioUrl" download>
        <button>Download Audio</button>
      </a>
    </div>
      </div>
    </div>
</template>

<script>
import axios from "axios"

export default {
  name: 'MainPage',
  data() {
    return {
      query: '',
      audioUrl: '',
      loading: false,
      //eslint-disable-next-line
      regExp: /^.*((youtu.be\/)|(v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))\??v?=?([^#\&\?]*).*/,
    };
  },
  methods: {
    
    async search() {
      this.loading = true;
      let extractVideo = (url) => {
        const match = url.match(this.regExp);
        if (match && match[7].length === 11) {
          return match[7];
        } else {
        // invalid YouTube URL
          return null;
        }
      }

      let videoID = extractVideo(this.query)
      await axios
        .get(`http://localhost:8080/api/audio/${videoID}`, {
          responseType: 'blob'
        })
        .then((response) => {
          this.loading = false;
          const url = window.URL.createObjectURL(new Blob([response.data], {type: 'audio/mp3'}));
          console.log(url);
          this.audioUrl = url;
        })
        .catch((error) => {
          this.loading = false;
          console.error(error);
        });
    }
  }
};
</script>

<style>
.form {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
input {
  padding: 0.5rem;
  font-size: 1.2rem;
  margin-right: 0.5rem;
}
button {
  padding: 0.5rem;
  font-size: 1.2rem;
  background-color: #007bff;
  color: #fff;
  border: none;
  cursor: pointer;
}
audio {
  width: 100%;
  margin-top: 1rem;
}
</style>