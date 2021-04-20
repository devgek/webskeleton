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
                data-target="#contactEditModal">{{$t("form.contact.list.buttonnew")}}</button>
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
                <tr :data-entityid="entity.ID" v-for="entity in Entities">
                    <td data-gkvval="entity.OrgType">{{ orgTypeDesc(entity.OrgType) }}</td>
                    <td data-gkvval="entity.Name">{{entity.Name}}</td>
                    <td data-gkvval="entity.NameExt">{{entity.NameExt}}</td>
                    <td data-gkvval="entity.ContactType">{{ contactTypeDesc(entity.ContactType)}}</td>
                    <td data-gkvval="entity.ID">{{entity.ID}}</td>
                    <td class="">
                        <div class="btn-group-sm" v-if="isAdminUser">
                            <button type="button" class="btn btn-sm btn-alt-primary gk-btn-edit" data-toggle="modal"
                                data-target="#contactEditModal">
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
</div>
<!-- END Page Container -->

`,
  created() {
    this.$store.dispatch("loadContacts");
  },
  methods: {
    contactTypeDesc(contactType) {
      return this.getContactTypes[contactType];
    },
    orgTypeDesc(orgType) {
      return this.getOrgTypes[orgType];
    },
  },
  computed: {
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
