const gekLayoutView = Vue.component("gek-layout-view", {
  template:
    /*html*/
 `
  <!-- Page Container -->
  <div id="page-container" class="page-header-dark main-content-xxx">
    <!-- Header -->
    <header id="page-header">
      <gek-header main-header="Go Webskeleton mit Vue frontend"></gek-header>
    </header>
    <!-- Main Container -->
    <main id="main-container">
      <gek-nav></gek-nav>
      <!-- Page Content -->
      <router-view></router-view>
      <!-- END Page Content -->
    </main>
    <!-- END Main Container -->
  </div>
  <!-- END Page Container -->
`,
    methods: {
    }
});
