Vue.component("gek-entity-edit-form-contact", {
  name: "gek-entity-edit-form-contact",
  props: {},
  template:
    /*html*/
    `<!-- EntityEditFormContact -->
  <div class="block-content font-size-sm">
    <div class="form-group">
        <label for="contactEditOrgType"
            class="col-form-label">{{$t("form.contact.edit.label.orgtype")}}</label>
        <select name="contactEditOrgType" class="form-control" id="contactEditOrgType" v-model="entityStores['Contact'].entityObject.OrgType">
          <option v-for="(option, key) in getOrgTypes" :value="key">{{ option }}</option>  
        </select>
    </div>
    <div class="form-group">
        <label for="contactEditName"
            class="col-form-label">{{$t("form.contact.edit.label.name")}}</label>
        <input name="contactEditName" class="form-control" id="contactEditName" v-model="entityStores['Contact'].entityObject.Name"
            autocomplete="new-password" />
    </div>
    <div class="form-group">
        <label for="contactEditNameExt"
            class="col-form-label">{{$t("form.contact.edit.label.nameext")}}</label>
        <input name="contactEditNameExt" class="form-control" id="contactEditNameExt" v-model="entityStores['Contact'].entityObject.NameExt"
            autocomplete="new-password" />
    </div>
    <div class="form-group">
        <label for="contactEditContactType"
            class="col-form-label">{{$t("form.contact.edit.label.contacttype")}}</label>
        <select name="contactEditContactType" class="form-control" id="contactEditContactType" v-model="entityStores['Contact'].entityObject.ContactType">
          <option v-for="(option, key) in getContactTypes" :value="key">{{ option }}</option>  
        </select>
    </div>
</div>
<!-- END EntityEditFormContact -->
`,
  data() {
    return {};
  },
  methods: {},
  computed: {
    ...Vuex.mapState(["entityStores"]),
    getOrgTypes() {
      return gkwebapp_T_OrgTypes;
    },
    getContactTypes() {
      return gkwebapp_T_ContactTypes;
    },
  },
});
