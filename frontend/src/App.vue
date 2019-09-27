<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/login">Login</router-link>

      <span v-if="isLoggedIn"> |
        <a href="" @click="logout">Logout</a>
      </span>
    </div>
    <router-view/>
  </div>
</template>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

<script>
import axios from "axios";

export default {
  computed: {
    isLoggedIn() {
      return this.$store.getters.isLoggedIn;
    }
  },

  methods: {
    logout() {
      // call logout mutator function and redirect to login page if successfull
      this.$store.dispatch('logout').then(() => {
        this.$router.push('/login');
      })
    }
  },

  created() {
    axios.interceptors.response.use(undefined, (err) => {
      return new Promise((response, reject) => {
        let res = err.response;

        if (res.status == 401 && res.config && !res.config.__isRetryRequest) {
          console.log("here");
          this.$store.dispatch('logout');
        }
        throw err;
      })
    })
  }
}
</script>