import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const AUTH_URL = 'http://localhost:8000/login';

export default new Vuex.Store({
  state: {
    accessToken: localStorage.getItem('token') || '',
    user: {},
    status: ''
  },
  mutations: {
    auth_request(state) {
      state.status = 'loading';
    },

    auth_success(state, token, user) {
      state.status = 'success';
      state.accessToken = token;
      state.user = user;
    },

    auth_error(state) {
      state.status = 'error';
    },

    logout(state) {
      state.status = '';
      state.accessToken = '';
    }

  },
  actions: {
    login({commit}, user) {
      return new Promise((resolve, reject) => {
        commit('auth_request');
        axios({url: AUTH_URL, data: user, method: 'POST'}).then(resp => {
          const token = resp.data.token;
          const user = resp.data.user;

          localStorage.setItem('token', token);
          axios.defaults.headers.common['Authorization'] = token

          commit('auth_success', token, user);
          
          resolve(resp);
        })
        .catch(err => {
          commit('auth_error');
          localStorage.removeItem('token');
          reject(err);
        })
      });
    },

    logout() {
      return new Promise((resolve, reject) => {
          commit('logout');
          localStorage.removeItem('token');

          delete axios.defaults.headers.common['Authorization'];
          resolve();
      });
    }
  },

  getters: {
    isLoggedIn: state => state.token,
    authStatus: state => state.status
  }
})
