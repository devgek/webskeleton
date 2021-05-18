const gekLoginView = Vue.component("gek-login-view", {
  props: {
    mainHeader: {
      type: String,
      default: "",
    },
    startPage: {
      type: String,
      default: "/noStartPage",
    }
  },
  template:
    /*html*/
    `    <div id="page-container">
      <main id="main-container">
          <div class="hero-static bg-white-95">
              <div class="content">
                  <div class="row justify-content-center">
                      <div class="col-md-8 col-lg-6 col-xl-4">
                          <div class="block block-themed block-fx-shadow mb-0">
                              <div class="block-header">
                                  <h3 class="block-title">Login</h3>
                              </div>
                              <div class="block-content">
                                  <div class="p-sm-3 px-lg-4 py-lg-5">
                                      <p><span
                                              class="font-w700 font-size-h5">{{ mainHeader }}</span>
                                      </p>
                                      <form class="" @submit.prevent="onSubmit">
                                          <div class="py-3">
                                              <div class="form-group">
                                                  <label for="login-username">Benutzer</label>
                                                  <input type="text" v-model="user" 
                                                      class="form-control form-control-alt form-control-lg"
                                                      id="login-username" name="userid">
                                              </div>
                                              <div class="form-group">
                                                  <label for="login-password">Passwort</label>
                                                  <input type="password" v-model="pass" 
                                                      class="form-control form-control-alt form-control-lg"
                                                      id="login-password" name="password">
                                              </div>
                                          </div>
                                          <div class="form-group row">
                                              <div class="col-md-6 col-xl-5">
                                                  <button type="submit" class="btn btn-block btn-primary">
                                                      <i class="fa fa-fw fa-sign-in-alt mr-1"></i>
                                                      Anmelden
                                                  </button>
                                              </div>
                                          </div>
                                          <div v-if="errorMessage" class="card-footer text-danger">
                                              {{ errorMessage }}
                                          </div>
                                      </form>
                                  </div>
                              </div>
                          </div>
                      </div>
                  </div>
              </div>

          </div>
      </main>
  </div>

      `,
  data() {
    return {
      user: "",
      pass: "",
      errorMessage: null
    };
  },
  methods: {
    onSubmit() {
      if (this.user === "" || this.pass === "") {
        this.errorMessage = "Username und Passwort müssen angegeben werden.";
        return;
      }

      this.errorMessage = "";

      this.$store
        .dispatch("login", {
          user: this.user,
          pass: this.pass,
        })
        .then(() => {
          this.$router.push({ name: this.startPage });
        })
        .catch((err) => {
          this.errorMessage = err.response.data;
        });
    },
  },
  computed: {
  }
});
