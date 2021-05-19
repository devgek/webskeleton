Vue.component("gek-entity-list-table-col-action", {
  props: {
    index: {
      type: Number,
      default: 0,
    },
    entity: {
      type: String,
      required: true,
    },
    entityName: {
      type: String,
      required: true,
    }
  },
  template:
    /*html*/
    `<!-- EntityListTableColAction -->
<td class="gk-col-edit" :data-index="index">
  <div class="btn-group-sm" v-if="isAdminUser">
      <button type="button" class="btn btn-sm btn-alt-primary gk-btn-edit" data-toggle="modal"
          :data-target="editModalIdRef" @click="SET_ENTITY_EDIT({entityName: entityName, entityIndex: index})">
          <i class="fa fa-fw fa-pencil-alt"></i>
      </button>
      <button type="button" class="btn btn-sm btn-alt-primary gk-btn-delete" data-toggle="modal"
          data-target="#confirmDeleteModal" @click="SET_ENTITY_EDIT({entityName: entityName, entityIndex: index})">
          <i class="fa fa-fw fa-times"></i>
      </button>
  </div>
  <div class="d-none" v-else="isAdminUser">
      &nbsp;
  </div>
</td>
<!-- END EntityListTableColAction -->
`,
  data() {
    return {
    };
  },
  methods: {
    ...Vuex.mapMutations([
      'SET_ENTITY_EDIT'
    ])
  },
  computed: {
    ...Vuex.mapGetters([
      'isAdminUser'
    ]),
    editModalIdRef() {
      return "#" + this.entity + "EditModal"
    }
  },
});
