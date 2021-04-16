// Create a new store instance.
const store = new Vuex.Store({
    state () {
      return {
        user: null
      }
    },
    mutations: {
      SET_USER_DATA (state, userData) {
        localStorage.setItem('user', JSON.stringify(userData))
        axios.defaults.headers.common['Authorization'] = `Bearer ${
          userData.token
        }`
        state.user = userData
      },
      LOGOUT () {
        localStorage.removeItem('user')
        location.reload()
      }
    },
    actions: {
      login ({ commit }, credentials) {
        return axios
          .post('//localhost:8080/loginapi', credentials)
          .then(({ data }) => {
            commit('SET_USER_DATA', data)
          })
      },
      logout ({ commit }) {
        commit('LOGOUT')
      }
    }
  })

