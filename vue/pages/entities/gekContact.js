const gekContactView = Vue.component("gek-contact", {
  template:
    /*html*/
`
<!-- Page Content -->
<div class="content content-full">
  <!-- entityEditDialog -->
  <gek-entity-edit-list entity="contact" entityName="Contact" />
  <!-- entityEditDialog -->
  <gek-entity-edit-dialog entity="contact" entityName="Contact" entityDesc="Kontakt" @entity-edit-save-contact="saveEntity({entityName:'Contact', entityDesc:'Kontakt'})"/>
  <!-- confirmDelete Dialog-->
  <gek-confirm-delete entity="contact" entityDesc="Kontakt" @entity-delete-confirm-contact="deleteEntity({entityName:'Contact', entityDesc:'Kontakt'})"/>
</div>
<!-- END Page Content -->
`,
  data() {
    return {};
  },
  created() {
    this.startEntityStore({
      entityName: "Contact",
      newEntityObjectFn: this.newEntityObject,
    });
  },
  methods: {
    ...Vuex.mapActions(["startEntityStore", "deleteEntity", "saveEntity"]),
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
  computed: {},
});
