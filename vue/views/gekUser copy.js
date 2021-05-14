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
    <div class="content content-full">
<!-- Your Block -->
<div class="block block-rounded">
    <div class="block-header block-header-default">
        <h3 class="block-title text-primary">{{$t("form.user.list.header")}}</h3>
    </div>
    <div class="block-content block-content-full font-size-sm">
        <div class="float-right mb-2">
            <button type="button" class="btn btn-outline-primary gk-btn-new" data-toggle="modal" data-target="#userEditModal" 
                 @click="prepareNew">{{$t("form.user.list.buttonnew")}}</button>
        </div>
        <div class="pb-3">&nbsp;</div>
    </div>
</div>
<div class="block block-rounded">
    <div class="block-content font-size-sm">
        <table class="table table-hover table-sm table-bordered table-striped gk-table js-dataTable-simple dataTable"
            id="userTable">
            <thead>
                <tr>
                    <th scope="col">Name</th>
                    <th scope="col">Passwort</th>
                    <th scope="col">Email</th>
                    <th scope="col">Benutzerrolle</th>
                    <th scope="col" class="w-5">{{$t("form.all.label.actions")}}</th>
                </tr>
            </thead>
            <tbody>
                <tr :data-entityid="entity.ID" :data-entityindex="index" v-for="(entity, index) in $store.getters.getEntityListByEntityName('User')" class="gk-row-edit">
                    <td :data-gkvval="entity.Name" class="gk-col-edit">{{entity.Name}}</td>
                    <td :data-gkvval="entity.Pass" class="gk-col-edit">********</td>
                    <td :data-gkvval="entity.Email" class="gk-col-edit">{{entity.Email}}</td>
                    <td :data-gkvval="entity.Role" class="gk-col-edit">{{ roleDesc(entity.Role) }}</td>
                    <td class="gk-col-edit">
                        <div class="btn-group-sm" v-if="$store.getters.isAdminUser">
                            <button type="button" class="btn btn-sm btn-alt-primary gk-btn-edit" data-toggle="modal"
                                data-target="#userEditModal" @click="prepareEntity(index)">
                                <i class="fa fa-fw fa-pencil-alt"></i>
                            </button>
                            <button type="button" class="btn btn-sm btn-alt-primary gk-btn-delete" data-toggle="modal"
                                data-target="#confirmDeleteModal" @click="prepareEntity(index)">
                                <i class="fa fa-fw fa-times"></i>
                            </button>
                        </div>
                        <div class="d-none" v-else="$store.getters.isAdminUser">
                            &nbsp;
                        </div>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>

</div>
<!-- END Your Block -->
    </div>
    <!-- END Page Content -->
  </main>
  <!-- END Main Container -->

  <!-- Modal Dialog-->
<div class="modal fade" id="userEditModal" data-backdrop="static" tabindex="-1" role="dialog"
aria-labelledby="userEditModalLabel" aria-hidden="true">
<div class="modal-dialog" role="document">
    <div class="modal-content">
        <div class="block block-themed block-transparent mb-0">
            <div class="block-header bg-primary">
                <h3 class="block-title" id="userEditModalLabel">{{ $store.state.entityStore.getEditHeader() }}</h3>
                <div class="block-options">
                    <button type="button" class="btn-block-option" data-dismiss="modal" aria-label="Close">
                        <i class="fa fa-fw fa-times"></i>
                    </button>
                </div>
            </div>
            <div class="block-content font-size-sm">
                <form>
                    <div class="form-group">
                        <label for="userEditName"
                            class="col-form-label">{{$t("form.user.edit.label.name")}}</label>
                        <input name="userEditName" class="form-control" id="userEditName" v-model="entityView.entityObject.Name" :readonly="!entityView.editNew"/>
                    </div>
                    <div class="form-group">
                        <label for="userEditPass"
                            class="col-form-label">{{$t("form.user.edit.label.pass")}}</label>
                        <input type="password" name="userEditPass" class="form-control" id="userEditPass" 
                            autocomplete="new-password" v-model="entityView.entityObject.Pass" :readonly="!entityView.editNew"/>
                    </div>
                    <div class="form-group">
                        <label for="userEditEmail"
                            class="col-form-label">{{$t("form.user.edit.label.email")}}</label>
                        <input name="userEditEmail" class="form-control" id="userEditEmail"  v-model="entityView.entityObject.Email"/>
                    </div>
                    <div class="form-group">
                        <label for="userEditRole"
                            class="col-form-label">{{$t("form.user.edit.label.role")}}</label>
                        <select name="userEditRole" class="form-control" id="userEditRole"  v-model="entityView.entityObject.Role">
                          <option v-for="(option, key) in getRoleTypes" :value="key">{{ option }}</option>  
                        </select>
                    </div>
                </form>
            </div>
            <div class="block-content block-content-full text-right border-top">
                <button type="button" class="btn btn-sm btn-light btn-back-app"
                    data-dismiss="modal">{{$t("form.all.btn.back")}}</button>
                <button type="button" class="btn btn-sm btn-primary btn-save-app" @click="$store.dispatch('saveEntity', {entityName:'User', entityObject: entityView.entityObject, editNew: entityView.editNew})" data-dismiss="modal">
                    <i class="fa fa-check mr-1"></i>{{$t("form.all.btn.save")}}
                </button>
            </div>
        </div>
    </div>
</div>
</div>
<!-- END Modal Dialog -->

<!-- confirmDelete Dialog-->
<gek-confirm-delete entity="user" entityName="Benutzer" @confirm-delete-user="$store.dispatch('deleteEntity', {entityName:'User', entityObject: entityView.entityObject, confirmed: $event})"/>

</div>
<!-- END Page Container -->
`,
  data() {
    return {
      entityView: new EntityView("User", this.newEntityObject(), this),
    };
  },
  created() {
    console.log("user created");
    this.$store.dispatch("loadEntities", {entityName: 'User'});
  },
  methods: {
    roleDesc(role) {
      return this.getRoleTypes[role];
    },
    prepareEntity(index) {
      //deep copy !
      var entityList = this.$store.getters.getEntityListByEntityName('User')
      this.entityView.entityObject = JSON.parse(JSON.stringify(entityList[index]));
      this.entityView.editNew = false;
    },
    prepareNew() {
      this.entityView.entityObject = this.newEntityObject();
      this.entityView.editNew = true;
    },
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
    getRoleTypes() {
      return gkwebapp_T_RoleTypes;
    },
    getContactTypes() {
      return gkwebapp_T_ContactTypes;
    },
    editHeader() {
      return this.entityView.getEditHeader();
    }
  },
});
