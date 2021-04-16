const app = new Vue({
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
  methods: {
    updateCart(id) {
      this.cart.push(id);
    },
  },
});
