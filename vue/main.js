const app = new Vue({
  i18n: i18n,
  store: store,
  router: router,
  el: "#app",
  data() {
    return {
      user: "",
      token: "",
      isAdmin: false,
    };
  },
  created() {
    const userString = localStorage.getItem('user')
    if (userString) {
      const userData = JSON.parse(userString)
      this.$store.commit('SET_USER_DATA', userData)
    }
    //
    axios.interceptors.response.use(
      response => response,
      error => {
        console.log(error.response)
        if (error.response.status === 401) {
          this.$router.push('/')
          this.$store.dispatch('logout')
        }
        return Promise.reject(error)
      }
    )
  },
});
