Vue.component("gek-entity-list-table-row-contact", {
  name: "gek-entity-list-table-row-contact",
  props: {
    entityObject: {
      type: Object,
      rquired: true
    },
    entityIndex: {
      type: Number,
      required: true
    }

  },
  template:
    /*html*/
    `<!-- EntityListTableHeaderContact -->
  <tr :data-entityid="entityObject.ID" :data-entityindex="entityIndex"class="gk-row-edit">
    <td :data-gkvval="entityObject.OrgType" class="gk-col-edit">{{ orgTypeDesc(entityObject.OrgType) }}</td>
    <td :data-gkvval="entityObject.Name" class="gk-col-edit">{{entityObject.Name}}</td>
    <td :data-gkvval="entityObject.NameExt" class="gk-col-edit">{{entityObject.NameExt}}</td>
    <td :data-gkvval="entityObject.ContactType" class="gk-col-edit">{{ contactTypeDesc(entityObject.ContactType)}}</td>
    <td :data-gkvval="entityObject.ID" class="gk-col-edit">{{entityObject.ID}}</td>

    <gek-entity-list-table-col-action entity="contact" entityName="Contact" :index="entityIndex"/>
  </tr>
<!-- END EntityListTableHeaderContact -->
`,
  data() {
    return {
    };
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
    getContactTypes() {
      return gkwebapp_T_ContactTypes;
    },
    getOrgTypes() {
      return gkwebapp_T_OrgTypes;
    },
  },
});
