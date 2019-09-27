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
      state.user = {};
    }

  },
  actions: {
    login({commit}, user) {
      return new Promise((resolve, reject) => {
        commit('auth_request');

        // login axios request
        axios.post(AUTH_URL, user).then(res => {
          let token = res.data.token;
          let user = res.data.user;
          
          // store token in local storage
          localStorage.setItem('token', token);
          axios.defaults.headers.common['Authorization'] = token

          // set state variables
          commit('auth_success', token, user);
          
          resolve(res);
        })
        .catch(err => {
          commit('auth_error');
          localStorage.removeItem('token');
          reject(err);
        })
      });
    },

    logout({commit}) {
      return new Promise((resolve, reject) => {
          // set state variables and remove token from local storage
          commit('logout');
          localStorage.removeItem('token');

          delete axios.defaults.headers.common['Authorization'];
          resolve();
      });
    }
  },

  getters: {
    isLoggedIn: state => !!state.accessToken,
    authStatus: state => state.status
  }
})
