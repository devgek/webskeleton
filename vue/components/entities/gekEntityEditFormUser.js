Vue.component("gek-entity-edit-form-user", {
  name: "gek-entity-edit-form-user",
  props: {
  },
  template:
    /*html*/
  `<!-- EntityEditFormUser -->
  <div class="block-content font-size-sm">
    <div class="form-group">
        <label for="userEditName"
            class="col-form-label">{{$t("form.user.edit.label.name")}}</label>
        <input name="userEditName" class="form-control" id="userEditName" v-model="entityStores['User'].entityObject.Name" :readonly="!entityStores['User'].editNew"/>
    </div>
    <div class="form-group">
        <label for="userEditPass"
            class="col-form-label">{{$t("form.user.edit.label.pass")}}</label>
        <input type="password" name="userEditPass" class="form-control" id="userEditPass" 
            autocomplete="new-password" v-model="entityStores['User'].entityObject.Pass" :readonly="!entityStores['User'].editNew"/>
    </div>
    <div class="form-group">
        <label for="userEditEmail"
            class="col-form-label">{{$t("form.user.edit.label.email")}}</label>
        <input name="userEditEmail" class="form-control" id="userEditEmail"  v-model="entityStores['User'].entityObject.Email"/>
    </div>
    <div class="form-group">
        <label for="userEditRole"
            class="col-form-label">{{$t("form.user.edit.label.role")}}</label>
        <select name="userEditRole" class="form-control" id="userEditRole"  v-model="entityStores['User'].entityObject.Role">
          <option v-for="(option, key) in getRoleTypes" :value="key">{{ option }}</option>  
        </select>
    </div>
</div>
<!-- END EntityEditFormUser -->
`,
  data() {
    return {
    };
  },
  methods: {
  },
  computed: {
    ...Vuex.mapState([
      'entityStores'
    ]),
    getRoleTypes() {
      return gkwebapp_T_RoleTypes;
    },
  },
});
