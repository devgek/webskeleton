const gekContactView = Vue.component("gek-contact", {
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
    <gek-entity-edit-list entity="contact" entityName="Contact" />
  <!-- END Page Content -->
</main>
<!-- END Main Container -->

<!-- entityEditDialog -->
<gek-entity-edit-dialog entity="contact" entityName="Contact" entityDesc="Kontakt" @entity-edit-save-contact="saveEntity({entityName:'Contact', entityDesc:'Kontakt'})"/>

<!-- confirmDelete Dialog-->
<gek-confirm-delete entity="contact" entityDesc="Kontakt" @entity-delete-confirm-contact="deleteEntity({entityName:'Contact', entityDesc:'Kontakt'})"/>

</div>
<!-- END Page Container -->
`,
data() {
  return {
  };
},
created() {
  this.startEntityStore({entityName: 'Contact', newEntityObjectFn: this.newEntityObject});
},
methods: {
  ...Vuex.mapActions([
    'startEntityStore',
    'deleteEntity',
    'saveEntity'
  ]),
  newEntityObject() {
      console.log("newEntityObject contact called");
      return {
        ID: 0,
        OrgType: 0,
        Name: "",
        NameExt: "",
        ContactType: 0,
      };
    },
  },
  computed: {
  },
});
