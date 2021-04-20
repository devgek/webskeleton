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
                data-target="#userEditModal">{{$t("form.user.list.buttonnew")}}</button>
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
                                data-target="#userEditModal">
                                <i class="fa fa-fw fa-pencil-alt"></i>
                            </button>
                            <button type="button" class="btn btn-sm btn-alt-primary gk-btn-delete" data-toggle="modal"
                                data-target="#confirmDeleteModal">
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
                        <input name="userEditName" class="form-control" id="userEditName" value="" />
                    </div>
                    <div class="form-group">
                        <label for="userEditPass"
                            class="col-form-label">{{$t("form.user.edit.label.pass")}}</label>
                        <input type="password" name="userEditPass" class="form-control" id="userEditPass" value=""
                            autocomplete="new-password" />
                    </div>
                    <div class="form-group">
                        <label for="userEditEmail"
                            class="col-form-label">{{$t("form.user.edit.label.email")}}</label>
                        <input name="userEditEmail" class="form-control" id="userEditEmail" value="" />
                    </div>
                    <div class="form-group">
                        <label for="userEditRole"
                            class="col-form-label">{{$t("form.user.edit.label.role")}}</label>
                        <select name="userEditRole" class="form-control" id="userEditRole" value="" v-model="userEditRole">
                          <option v-for="(option, key) in getRoleTypes" :value="key">{{ option }}</option>  
                        </select>
                    </div>
                </form>
            </div>
            <div class="block-content block-content-full text-right border-top">
                <button type="button" class="btn btn-sm btn-light btn-back-app"
                    data-dismiss="modal">{{$t("form.all.btn.abort")}}</button>
                <button type="button" class="btn btn-sm btn-primary btn-save-app">
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
      userEditRole: 0,
    };
  },
  created() {
    this.$store.dispatch("loadUsers");
  },
  methods: {
    roleDesc(role) {
      return this.getRoleTypes[role];
    },
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
  },
});

//create GKTable without inline editing
var userTable = new GKEntityTable("user");

activeGKEntityTable.prepareEditDialog = function () {
  if (activeGKEntityTable.isEditNew()) {
    $("#userEditModalLabel").html("{{$headerNew}}");
    $("#userEditName").prop("readonly", false);
    $("#userEditPass").prop("readonly", false);
  } else {
    $("#userEditModalLabel").html("{{$headerEdit}}");
    $("#userEditName").prop("readonly", true);
    $("#userEditPass").prop("readonly", true);
  }

  $("#userEditName").val(activeGKEntityTable.editRowData[0]);
  $("#userEditPass").val(activeGKEntityTable.editRowDataHidden[1]);
  $("#userEditEmail").val(activeGKEntityTable.editRowDataHidden[2]);
  $("#userEditRole").val(activeGKEntityTable.editRowDataHidden[3]);
};
activeGKEntityTable.prepareSendRowData = function () {
  var sendParams = [];
  sendParams["gkvObjId"] = activeGKEntityTable.getEditRowKey();
  sendParams["gkvName"] = $("#userEditName").val();
  sendParams["gkvPass"] = $("#userEditPass").val();
  sendParams["gkvEmail"] = $("#userEditEmail").val();
  sendParams["gkvRole"] = $("#userEditRole").val();

  return sendParams;
};
activeGKEntityTable.getRowDataFromEntity = function (data) {
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
activeGKEntityTable.getRowDataHiddenFromEntity = function (data) {
  var rowDataHidden = [];
  rowDataHidden.push(
    data.EntityObject.Name,
    data.EntityObject.Pass,
    data.EntityObject.Email,
    data.EntityObject.Role
  );
  return rowDataHidden;
};
