const gekUserView = Vue.component("gek-user", {
  template:
    /*html*/
    `
  <div id="page-container" class="page-header-dark main-content-xxx">
  <!-- Header -->
  <header id="page-header">
    <gek-header main-header="Go Webskeleton mit Vue frontend"></gek-header>
  </header>
  <!-- Main Container -->
  <main id="main-container">
    <gek-nav></gek-nav>
    <!-- Page Content -->
      <!-- entityEditDialog -->
      <gek-entity-edit-list entity="user" entityName="Benutzer" />
    <!-- END Page Content -->
  </main>
  <!-- END Main Container -->

<!-- entityEditDialog -->
<gek-entity-edit-dialog entity="user" entityName="Benutzer" @entity-edit-save-user="saveEntity({entityName:'User'})"/>

<!-- confirmDelete Dialog-->
<gek-confirm-delete entity="user" entityName="Benutzer" @entity-delete-confirm-user="deleteEntity({entityName:'User'})"/>

</div>
<!-- END Page Container -->
`,
  data() {
    return {
    };
  },
  created() {
    console.log("user created");
    this.startEntityStore({entityName: 'User', newEntityObjectFn: this.newEntityObject});
  },
  methods: {
    ...Vuex.mapActions([
      'startEntityStore',
      'deleteEntity',
      'saveEntity'
    ]),
    newEntityObject() {
      console.log("newEntityObject user called");
      return {
        ID: 0,
        Name: "",
        Pass: "",
        Email: "",
        Role: 0,
      };
        }
  },
  computed: {
  },
});
