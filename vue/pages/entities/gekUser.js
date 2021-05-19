const gekUserView = Vue.component("gek-user", {
  template:
    /*html*/
 `
<!-- Page Content -->
<div class="content content-full">
  <!-- entityEditDialog -->
  <gek-entity-edit-list entity="user" entityName="User" />
  <!-- entityEditDialog -->
  <gek-entity-edit-dialog entity="user" entityName="User" entityDesc="Benutzer" @entity-edit-save-user="saveEntity({entityName:'User', entityDesc:'Benutzer'})"/>
  <!-- confirmDelete Dialog-->
  <gek-confirm-delete entity="user" entityDesc="Benutzer" @entity-delete-confirm-user="deleteEntity({entityName:'User', entityDesc:'Benutzer'})"/>
</div>
<!-- END Page Content -->
`,
  data() {
    return {
    };
  },
  created() {
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
