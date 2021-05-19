Vue.component("gek-entity-edit-list", {
  props: {
    entity: {
      type: String,
      required: true,
    },
    entityName: {
      type: String,
      required: true,
    },
  },
  template:
    /*html*/
  `
  <div>
    <!-- Your Block -->
    <div class="block block-rounded">
    <div class="block-header block-header-default">
        <h3 class="block-title text-primary">{{$t("form." + entity + ".list.header")}}</h3>
    </div>
    <div class="block-content block-content-full font-size-sm">
        <div class="float-right mb-2">
            <button type="button" class="btn btn-outline-primary gk-btn-new" data-toggle="modal" :data-target="editModalIdRef" 
                @click="SET_ENTITY_NEW({entityName: entityName, entityDesc: 'Benutzer'})">{{$t("form." + entity + ".list.buttonnew")}}</button>
        </div>
        <div class="pb-3">&nbsp;</div>
    </div>
  </div>
  <div class="block block-rounded">
    <div class="block-content font-size-sm">
        <table class="table table-hover table-sm table-bordered table-striped gk-table js-dataTable-simple dataTable">
            <thead>
              <component :is="entityListTableHeaderComponent"></component>
            </thead>
            <tbody>
              <component :is="entityListTableRowComponent" v-for="(entityObject, index) in getEntityListByEntityName(entityName)" :entityObject="entityObject" :entityIndex="index"></component>
            </tbody>
        </table>
    </div>
  </div>
  <!-- END Your Block -->
</div>
`,
  data() {
    return {
      entityListTableHeaderComponent: "gek-entity-list-table-header-" + this.entity,
      entityListTableRowComponent: "gek-entity-list-table-row-" + this.entity,
    };
  },
  methods: {
    ...Vuex.mapMutations([
      'SET_ENTITY_NEW'
    ])
  },
  computed: {
    ...Vuex.mapState([
      'entityStores'
    ]),
    ...Vuex.mapGetters([
      'isAdminUser',
      'getEntityListByEntityName',
      'getUser'
    ]),
    editModalIdRef() {
      return "#" + this.entity + "EditModal"
    }
  },
});
