Vue.component("gek-error-message", {
  props: {},
  template:
    /*html*/
    `<div class="toast bg-warning" role="alert" aria-live="assertive" aria-atomic="true"
    data-delay="3000" data-toastid="errorMessageToast" >
    <div class="toast-header">
        <span class="gk-toast-text" v-if="message">{{message.msg}}</span>
        <button type="button" class="ml-2 mb-1 close" data-dismiss="toast" aria-label="Close">
            <span aria-hidden="true">&times;</span>
        </button>
    </div>
  </div>
`,
  data() {
    return {};
  },
  created() {
    this.unsubscribe = this.$store.subscribe((mutation, state) => {
      if (mutation.type === "SET_MESSAGE") {
        console.log(`catching message from store: ${state.message.msg}`);

        const $toast = $(".toast[data-toastid='errorMessageToast']");
        $toast.toast("show");
      }
    });
  },
  beforeDestroy() {
    this.unsubscribe();
  },
  methods: {
  },
  computed: {
    message() {
      return this.$store.state.message;
    },
  },
});
