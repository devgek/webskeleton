const gekHomeView = Vue.component("gek-home", {
  template:
    /*html*/
    `  <div id="page-container" class="page-header-dark main-content-xxx">
    <!-- Header -->
    <header id="page-header">
      <gek-header main-header="Go Webskeleton mit Vue frontend"></gek-header>
    </header>
    <!-- Main Container -->
    <main id="main-container">
      <!-- Page Content -->
      <div class="content content-full">
        <!-- Your Block -->
        Home view
        <router-link to="/page1" class="font-w700 font-size-h5">Seite 1</router-link>
        <router-link to="/page2" class="font-w700 font-size-h5">Seite 2</router-link>
        <button type="button">Count</button>
        <button type="button" class="logoutButton" @click="logout">Logout</button>
        <!-- END Your Block -->
      </div>
      <!-- END Page Content -->
    </main>
    <!-- END Main Container -->
  </div>
  <!-- END Page Container -->`,
  methods: {
    logout () {
      this.$store.dispatch('logout')
    }
  }
});
