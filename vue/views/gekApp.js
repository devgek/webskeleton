const gekAppView = Vue.component("gek-app", {
  template:
    /*html*/
    `<router-view @login-submit="doLogin"></router-view>`,
  methods: {
    doLogin(loginData) {
      alert('login in gek-app with ' + loginData.user + " " + loginData.pass)
      this.$store
      .dispatch('loginapi', {
        user: loginData.user,
        password: loginData.pass
      })
      .then(() => { this.$router.push({ name: 'Page1' }) })
      .catch(err => { this.status = err.response.status })
    }
  }
});
