var forceKey = 17;
var editName = "";
var editPass = "";
var editEmail = "";
var editRole = 0;

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
            <button type="button" class="btn btn-outline-primary gk-btn-new" data-toggle="modal"
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
                <tr :data-entityid="entity.ID" v-for="entity in Entities">
                    <td :data-gkvval="entity.Name">{{entity.Name}}</td>
                    <td :data-gkvval="entity.Pass">********</td>
                    <td :data-gkvval="entity.Email">{{entity.Email}}</td>
                    <td :data-gkvval="entity.Role">{{ roleDesc(entity.Role) }}</td>
                    <td class="">
                        <div class="btn-group-sm" v-if="isAdminUser">
                            <button type="button" class="btn btn-sm btn-alt-primary gk-btn-edit" data-toggle="modal"
                                data-target="#userEditModal" @click="prepareEdit">
                                <i class="fa fa-fw fa-pencil-alt"></i>
                            </button>
                            <button type="button" class="btn btn-sm btn-alt-primary gk-btn-delete" data-toggle="modal"
                                data-target="#confirmDeleteModal" @click="prepareDelete">
                                <i class="fa fa-fw fa-times"></i>
                            </button>
                        </div>
                        <div class="d-none" v-else="isAdminUser">
                            &nbsp;
                        </div>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
    <div class="block-content font-size-sm">
    <form>
        <div class="form-group">
            <label for="userTestName"
                class="col-form-label">TestName</label>
            <input name="userTestName" class="form-control" id="userTestName" :value="name" />
        </div>
    </form>
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
                <h3 class="block-title" id="userEditModalLabel">Modal Title</h3>
                <div class="toast bg-warning" role="alert" aria-live="assertive" aria-atomic="true"
                    data-delay="3000" data-tableId="userTable">
                    <div class="toast-header">
                        <span class="gk-toast-text">Toast Text</span>
                        <button type="button" class="ml-2 mb-1 close" data-dismiss="toast" aria-label="Close">
                            <span aria-hidden="true">&times;</span>
                        </button>
                    </div>
                </div>
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
                        <input name="userEditName" class="form-control" id="userEditName" :value="editName"/>
                    </div>
                    <div class="form-group">
                        <label for="userEditPass"
                            class="col-form-label">{{$t("form.user.edit.label.pass")}}</label>
                        <input type="password" name="userEditPass" class="form-control" id="userEditPass" 
                            autocomplete="new-password" :value="editPass"/>
                    </div>
                    <div class="form-group">
                        <label for="userEditEmail"
                            class="col-form-label">{{$t("form.user.edit.label.email")}}</label>
                        <input name="userEditEmail" class="form-control" id="userEditEmail"  :value="editEmail"/>
                    </div>
                    <div class="form-group">
                        <label for="userEditRole"
                            class="col-form-label">{{$t("form.user.edit.label.role")}}</label>
                        <select name="userEditRole" class="form-control" id="userEditRole"  :value="editRole">
                          <option v-for="(option, key) in getRoleTypes" :value="key">{{ option }}</option>  
                        </select>
                    </div>
                </form>
            </div>
            <div class="block-content block-content-full text-right border-top">
                <button type="button" class="btn btn-sm btn-light btn-back-app"
                    data-dismiss="modal">{{$t("form.all.btn.abort")}}</button>
                <button type="button" class="btn btn-sm btn-primary btn-save-app" @click="doSave">
                    <i class="fa fa-check mr-1"></i>{{$t("form.all.btn.save")}}
                </button>
            </div>
        </div>
    </div>
</div>
</div>
<!-- END Modal Dialog -->
</div>
<!-- END Page Container -->
`,
  data() {
    return {
      editName,
      editPass,
      editEmail,
      editRole
    };
  },
  created() {
    this.$store.dispatch("loadUsers");
  },
  updated() {
    console.log("user updated");
    // console.log(this.$el.innerHTML);
    userTable.initialize();
  },
  methods: {
    roleDesc(role) {
      return this.getRoleTypes[role];
    },
    doSave() {
      console.log("doSave:" + this.entityObject.Name + "," + this.entityObject.Pass + "," + this.entityObject.Email + "," + this.entityObject.Role)
      this.$store.state.editUser = this.entityObject
    },
    prepareNew() {
      this.$store.commit('SET_EDIT_USER_NAME', "new")
      this.$store.commit('SET_EDIT_USER_PASS', "xxx")
    },
    prepareEdit() {
      this.$store.commit('SET_EDIT_USER_NAME', this.editName)
      this.$store.commit('SET_EDIT_USER_PASS', this.editPass)
      this.forceKey++;
    },
    prepareDelete() {
    },
    storeName(event) {
      this.$store.commit('SET_EDIT_USER_NAME', "krxmxr")
    }
  },
  computed: {
    getRoleTypes() {
      return gkwebapp_T_RoleTypes;
    },
    getContactTypes() {
      return gkwebapp_T_ContactTypes;
    },
    isAdminUser() {
      return this.$store.state.isAdmin;
    },
    Entities() {
      return this.$store.state.users;
    },
    name: {
      get () {
        return this.$store.state.editUser.Name
      },
      set(value) {
        this.$store.commit('SET_EDIT_USER_NAME', value)
      }
    },
    pass: {
      get () {
        return this.$store.state.editUser.Pass
      },
      set(value) {
        this.$store.commit('SET_EDIT_USER_PASS', value)
      }
    },
    email: {
      get () {
        return this.$store.state.editUser.Email
      },
      set(value) {
        this.$store.commit('SET_EDIT_USER_EMAIL', value)
      }
    },
    role: {
      get () {
        return this.$store.state.editUser.Role
      },
      set(value) {
        this.$store.commit('SET_EDIT_USER_ROLE', value)
      }
    }
  
  },
});

//create GKTable without inline editing
var userTable = new GKEntityTable("user");

userTable.prepareEditDialog = function () {
  if (activeGKEntityTable.isEditNew()) {
    $("#userEditModalLabel").html("Benutzer neu anlegen");
    $("#userEditName").prop("readonly", false);
    $("#userEditPass").prop("readonly", false);
  } else {
    $("#userEditModalLabel").html("Benutzer ändern");
    $("#userEditName").prop("readonly", true);
    $("#userEditPass").prop("readonly", true);
  }

    console.log("rowData:" + activeGKEntityTable.editRowData[0] + "," + activeGKEntityTable.editRowData[1]);
    console.log("rowDataHidden:" + activeGKEntityTable.editRowDataHidden[0] + "," + activeGKEntityTable.editRowDataHidden[1]);
    console.log("store:" + JSON.stringify(store.state.editUser))

    editName = activeGKEntityTable.editRowData[0];
    editPass = activeGKEntityTable.editRowDataHidden[1];
    editEmail = activeGKEntityTable.editRowDataHidden[2];
    editRole = activeGKEntityTable.editRowDataHidden[3];

    forceKey++;
};
userTable.prepareSendRowData = function () {
  console.log("getUserView:prepareSendRowData:" + JSON.stringify(store.state.editUser))
  return store.state.editUser
};
userTable.getRowDataFromEntity = function (data) {
  var rowData = [];
  var roleName = gkwebapp_T_RoleTypes[data.EntityObject.Role];
  rowData.push(
    data.EntityObject.Name,
    "********",
    data.EntityObject.Email,
    roleName
  );
  return rowData;
};
userTable.getRowDataHiddenFromEntity = function (data) {
  var rowDataHidden = [];
  rowDataHidden.push(
    data.EntityObject.Name,
    data.EntityObject.Pass,
    data.EntityObject.Email,
    data.EntityObject.Role
  );
  return rowDataHidden;
};
