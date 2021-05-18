Vue.component("gek-entity-list-table-row-user", {
  name: "gek-entity-list-table-row-user",
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
    `<!-- EntityListTableHeaderUser -->
  <tr :data-entityid="entityObject.ID" :data-entityindex="entityIndex"class="gk-row-edit">
    <td :data-gkvval="entityObject.Name" class="gk-col-edit">{{entityObject.Name}}</td>
    <td :data-gkvval="entityObject.Pass" class="gk-col-edit">********</td>
    <td :data-gkvval="entityObject.Email" class="gk-col-edit">{{entityObject.Email}}</td>
    <td :data-gkvval="entityObject.Role" class="gk-col-edit">{{ roleDesc(entityObject.Role) }}</td>

    <gek-entity-list-table-col-action entity="user" entityName="User" :index="entityIndex"/>
  </tr>
<!-- END EntityListTableHeaderUser -->
`,
  data() {
    return {
    };
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
  },
});
