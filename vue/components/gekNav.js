Vue.component("gek-nav", {
  template:
    /*html*/
    `            <!-- Navigation -->
  <div class="bg-white">
      <div class="content py-3">
          <!-- Main Navigation -->
          <div id="main-navigation" class="d-none d-lg-block mt-2 mt-lg-0">
              <ul class="nav-main nav-main-horizontal nav-main-hover">
                  <li class="nav-main-item">
                      <a class="nav-main-link nav-main-link-submenu" data-toggle="submenu"
                          aria-haspopup="true" aria-expanded="false" >
                          <i class="nav-main-link-icon si si-speedometer"></i>
                          <span
                              class="nav-main-link-name">{{$t("nav.pages.header")}}</span>
                      </a>
                      <ul class="nav-main-submenu">
                          <li class="nav-main-item">
                                <span class="nav-main-link-name"><router-link to="/page1" class="nav-main-link">{{$t("nav.pages.page1")}}</router-link></span>
                          </li>
                      </ul>
                  </li>
                  <li class="nav-main-heading">Heading</li>
                  <li class="nav-main-item">
                      <a class="nav-main-link nav-main-link-submenu" data-toggle="submenu"
                          aria-haspopup="true" aria-expanded="false" >
                          <i class="nav-main-link-icon si si-settings"></i>
                          <span
                              class="nav-main-link-name">{{$t("nav.admin.header")}}</span>
                      </a>
                      <ul class="nav-main-submenu">
                          <li class="nav-main-item">
                              <a class="nav-main-link" href="/entitylistuser">
                                  <span
                                      class="nav-main-link-name"><router-link to="/admin/user" class="nav-main-link">{{$t("nav.admin.user")}}</router-link></span>
                              </a>
                          </li>
                          <li class="nav-main-item">
                              <a class="nav-main-link" href="/entitylistcontact">
                                  <span
                                      class="nav-main-link-name"><router-link to="/admin/contact" class="nav-main-link">{{$t("nav.admin.contact")}}</router-link></span>
                              </a>
                          </li>
                      </ul>
                  </li>
              </ul>
              <div class="float-right error-message-container">
                <gek-error-message/>
              </div>
          </div>
          <!-- END Main Navigation -->

      </div>
  </div>
  <!-- END Navigation -->
`,
  data() {
      return {}
  },
  methods: {
  },
  computed: {
  },
});
