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
    <div class="content content-full">
<!-- Your Block -->
<div class="block block-rounded">
    <div class="block-header">
        <h3 class="block-title text-primary">{{$t("form.contact.list.header")}}</h3>
        <span class="float-right">
            <button type="button" class="btn btn-outline-primary gk-btn-new" data-toggle="modal"
                data-target="#contactEditModal" @click="prepareNew($event.target)">{{$t("form.contact.list.buttonnew")}}</button>
        </span>
    </div>
    <div class="block-content font-size-sm">
        <table class="table table-hover table-sm table-bordered table-striped gk-table js-dataTable-simple dataTable"
            id="contactTable">
            <thead>
                <tr>
                    <th scope="col">{{$t("form.contact.list.orgtype")}}</th>
                    <th scope="col">{{$t("form.contact.list.name")}}</th>
                    <th scope="col">{{$t("form.contact.list.nameext")}}</th>
                    <th scope="col">{{$t("form.contact.list.contacttype")}}</th>
                    <th scope="col">{{$t("form.contact.list.id")}}</th>
                    <th scope="col" class="w-5">{{$t("form.all.label.actions")}}</th>
                </tr>
            </thead>
            <tbody>
                <tr :data-entityid="entity.ID" v-for="entity in Entities" class="gk-row-edit">
                    <td data-gkvval="entity.OrgType" class="gk-col-edit">{{ orgTypeDesc(entity.OrgType) }}</td>
                    <td data-gkvval="entity.Name" class="gk-col-edit">{{entity.Name}}</td>
                    <td data-gkvval="entity.NameExt" class="gk-col-edit">{{entity.NameExt}}</td>
                    <td data-gkvval="entity.ContactType" class="gk-col-edit">{{ contactTypeDesc(entity.ContactType)}}</td>
                    <td data-gkvval="entity.ID" class="gk-col-edit">{{entity.ID}}</td>
                    <td class="gk-col-edit">
                        <div class="btn-group-sm" v-if="isAdminUser">
                            <button type="button" class="btn btn-sm btn-alt-primary gk-btn-edit" data-toggle="modal"
                                data-target="#contactEditModal" @click="prepareEdit($event.target)">
                                <i class="fa fa-fw fa-pencil-alt"></i>
                            </button>
                            <button type="button" class="btn btn-sm btn-alt-primary gk-btn-delete" data-toggle="modal"
                                data-target="#confirmDeleteModal" @click="prepareDelete($event.target)">
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
<div class="modal fade" id="contactEditModal" data-backdrop="static" tabindex="-1" role="dialog"
aria-labelledby="contactEditModalLabel" aria-hidden="true">
<div class="modal-dialog" role="document">
    <div class="modal-content">
        <div class="block block-themed block-transparent mb-0">
            <div class="block-header bg-primary">
                <h3 class="block-title" id="contactEditModalLabel">{{header}}</h3>
                <div class="block-options">
                    <button type="button" class="btn-block-option" data-dismiss="modal" aria-label="Close">
                        <i class="fa fa-fw fa-times"></i>
                    </button>
                </div>
            </div>
            <div class="block-content font-size-sm">
                <div class="form-group">
                    <label for="contactEditOrgType"
                        class="col-form-label">{{$t("form.contact.edit.label.orgtype")}}</label>
                    <select name="contactEditOrgType" class="form-control" id="contactEditOrgType" v-model="entityObject.orgType">
                      <option v-for="(option, key) in getOrgTypes" :value="key">{{ option }}</option>  
                    </select>
                </div>
                <div class="form-group">
                    <label for="contactEditName"
                        class="col-form-label">{{$t("form.contact.edit.label.name")}}</label>
                    <input name="contactEditName" class="form-control" id="contactEditName" v-model="entityObject.name"
                        autocomplete="new-password" />
                </div>
                <div class="form-group">
                    <label for="contactEditNameExt"
                        class="col-form-label">{{$t("form.contact.edit.label.nameext")}}</label>
                    <input name="contactEditNameExt" class="form-control" id="contactEditNameExt" v-model="entityObject.nameExt"
                        autocomplete="new-password" />
                </div>
                <div class="form-group">
                    <label for="contactEditContactType"
                        class="col-form-label">{{$t("form.contact.edit.label.ContactType")}}</label>
                    <select name="contactEditContactType" class="form-control" id="contactEditContactType" v-model="entityObject.contactType">
                      <option v-for="(option, key) in getContactTypes" :value="key">{{ option }}</option>  
                    </select>
                </div>
            </div>
            <div class="block-content block-content-full text-right border-top">
                <button type="button" class="btn btn-sm btn-light btn-back-app"
                    data-dismiss="modal" @click="">{{$t("form.all.btn.back")}}</button>
                <button type="button" class="btn btn-sm btn-primary btn-save-app" @click="doSave" data-dismiss="modal">
                    <i class="fa fa-check mr-1"></i>{{$t("form.all.btn.save")}}
                </button>
            </div>
        </div>
    </div>
</div>
</div>

</div>
<!-- END Page Container -->

`,
data() {
  return {
    contactTable: new GKEntityTable("contact"),
    entityObject: this.newEntityObject(),
  };
},
created() {
  console.log("contact created");
  this.$store.dispatch("loadContacts");

  this.contactTable.prepareEditDialog = function () {
  };

  this.contactTable.getRowDataFromEntity = function (data) {
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
  this.contactTable.getRowDataHiddenFromEntity = function (data) {
    var rowDataHidden = [];
    rowDataHidden.push(
      data.EntityObject.Name,
      data.EntityObject.Pass,
      data.EntityObject.Email,
      data.EntityObject.Role
    );
    return rowDataHidden;
  };

  this.editName = this.editName + "created";
},
updated() {
  console.log("contact updated");
  this.contactTable.initialize();
},
mounted() {
  console.log("contact mounted");
},
methods: {
    contactTypeDesc(contactType) {
      return this.getContactTypes[contactType];
    },
    orgTypeDesc(orgType) {
      return this.getOrgTypes[orgType];
    },
    doSave() {
     if (this.contactTable.isEditNew()) {
        this.$store.dispatch("createContact", this.entityObject);
      } else {
        this.$store.dispatch("updateContact", this.entityObject);
      }
    },
    doDelete(confirmed) {
      if (confirmed) {
        this.$store.dispatch("deleteContact", this.entityObject);
      }
    },
    prepareNew(eventTarget) {
      console.log("prepareNew contact");

      this.contactTable.onStartRowEditing(eventTarget);
      this.entityObject = this.newEntityObject();
    },
    prepareEdit(eventTarget) {
      console.log("prepareEdit contact");

      this.contactTable.onStartRowEditing(eventTarget);

      this.entityObject.id = parseInt(this.contactTable.editRowKey);

      this.entityObject.orgType = parseInt(this.contactTable.editRowDataHidden[0]);
      this.entityObject.name = this.contactTable.editRowData[1];
      this.entityObject.nameExt = this.contactTable.editRowData[2];
      this.entityObject.contactType = parseInt(this.contactTable.editRowDataHidden[3]);
    },
    prepareDelete(eventTarget) {
      console.log("prepareDelete contact");

      this.contactTable.onStartRowEditing(eventTarget);

      this.entityObject.id = parseInt(this.contactTable.editRowKey);
    },
    newEntityObject() {
      console.log("newEntityObject contact called");
      return {
        id: 0,
        orgType: 0,
        name: "",
        nameExt: "",
        contactType: 0,
      };
    },
  },
  computed: {
    header() {
      if (this.contactTable.isEditNew()) {
        return 'Kontakt neu anlegen'
      }
      else {
        return 'Kontakt ändern';
      }
    },
    isEditEdit() {
      return !this.contactTable.isEditNew()
    },
    getOrgTypes() {
      return gkwebapp_T_OrgTypes;
    },
    getContactTypes() {
      return gkwebapp_T_ContactTypes;
    },
    isAdminUser() {
      return this.$store.state.isAdmin;
    },
    Entities() {
      return this.$store.state.contacts
    }
  },
});
