const gekAppView = Vue.component("gek-app", {
  template:
    /*html*/
    `<router-view></router-view>`,
  methods: {
  },
  created() {
    this.unsubscribe = this.$store.subscribe((mutation, state) => {
      if (mutation.type === "LOGOUT") {
        console.log("logout catched in gek-app");
        this.$router.push({ name: "Login" })
      }
    });
  },
  beforeDestroy() {
    console.log("gek-app destroyed");
    this.unsubscribe();
  },

});
