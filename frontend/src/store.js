import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    accessToken: localStorage.getItem('access_token') || '',
    status: {}
  },
  mutations: {
    request(state) {
      state.status = 'loading';
    },

    sucess(state, token) {
      state.status = 'success';
      state.accessToken = token;
    },

    error(state) {
      state.status = 'error';
    }

  },
  actions: {
    login({commit}, user) {
      return new Promise((resolve, reject) => {
        commit(login);
        axios({url: 'auth', data: user, method: 'POST'}).then(resp => {
          const token = resp.data.token;
          localStorage.setItem('user-token', token);
          commit(success, token);
          
          dispatch(request);
          resolve(resp);
        })
        .catch(err => {
          commit(error, err);
          localStorage.removeItem('user-token');
          reject(err);
        })
      });
    },

    logout() {
      this.$store.dispatch()
    }
  }
})
