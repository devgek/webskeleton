const gekLayoutView = Vue.component("gek-layout-view", {
  template:
    /*html*/
    `    <router-view @login-submit="doxx"></router-view>
    `,
    methods: {
      doxx(loginData) {
        alert("login!!!")
      }
    }
});
