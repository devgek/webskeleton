Vue.component("gek-header", {
  props: {
    mainHeader: {
      type: String,
      required: true,
    },
  },
  template:
    /*html*/
    `        <!-- Header -->
  <header id="page-header">
      <!-- Header Content -->
      <div class="content-header">
          <!-- Left Section -->
          <div class="d-flex align-items-center">
              <!-- Logo -->
                <span class="font-w700 font-size-h5 text-dual"><router-link to="/home" class="font-w700 font-size-h5">{{ mainHeader }} - {{ $store.state.user.name }}</router-link></span>
              <!-- END Logo -->

          </div>
          <!-- END Left Section -->

          <!-- Right Section -->
          <div class="d-flex align-items-center">
              <!-- User Dropdown -->
              <div class="dropdown d-inline-block ml-2">
                  <button v-if="user" type="button" class="btn btn-sm btn-dual" id="page-header-user-dropdown"
                      data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <img class="rounded" src="/assets/media/avatars/avatar10.jpg" alt="Header Avatar"
                          style="width: 18px;">
                      <span class="d-none d-sm-inline-block ml-1">{{ user.name }}</span>
                      <i class="fa fa-fw fa-angle-down d-none d-sm-inline-block"></i>
                  </button>
                  <div class="dropdown-menu dropdown-menu-right p-0 border-0 font-size-sm"
                      aria-labelledby="page-header-user-dropdown">
                      <div class="p-3 text-center bg-primary">
                          <img class="img-avatar img-avatar48 img-avatar-thumb"
                              src="/assets/media/avatars/avatar10.jpg" alt="">
                      </div>
                      <div class="p-2">
                          <h5 class="dropdown-header text-uppercase">Aktionen</h5>
                          <div class="dropdown-item d-flex align-items-center justify-content-between" @click="logout" style="cursor:pointer;">
                              <span>Logout</span>
                              <i class="si si-logout ml-1"></i>
                          </div>
                      </div>
                  </div>
              </div>
              <!-- END User Dropdown -->
          </div>
          <!-- END Right Section -->
      </div>
      <!-- END Header Content -->

      <!-- Header Loader -->
      <!-- Please check out the Loaders page under Components category to see examples of showing/hiding it -->
      <div id="page-header-loader" class="overlay-header bg-primary-lighter">
          <div class="content-header">
              <div class="w-100 text-center">
                  <i class="fa fa-fw fa-circle-notch fa-spin text-primary"></i>
              </div>
          </div>
      </div>
      <!-- END Header Loader -->
  </header>
  <!-- END Header -->
`,
  data() {
    return {};
  },
  methods: {
    logout() {
      this.$store.dispatch("logout");
    },
  },
  computed: {
    user() {
      return this.$store.getters.getUser;
    },
  },
});
