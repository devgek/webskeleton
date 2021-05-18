Vue.component("gek-entity-edit-dialog", {
  props: {
    entity: {
      type: String,
      required: true,
    },
    entityName: {
      type: String,
      required: true,
    },
    entityDesc: {
      type: String,
      required: true,
    }
  },
  template:
    /*html*/
  `<!-- Modal Dialog-->
    <div class="modal fade" :id="editModalId" data-backdrop="static" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="block block-themed block-transparent mb-0">
                <div class="block-header bg-primary">
                    <h3 class="block-title">{{ entityStores[entityName].getEditHeader(entityDesc) }}</h3>
                    <div class="block-options">
                        <button type="button" class="btn-block-option" data-dismiss="modal" aria-label="Close">
                            <i class="fa fa-fw fa-times"></i>
                        </button>
                    </div>
                </div>
                <!-- include the edit fields here -->
                <component :is="editFormComponent"></component>

                <div class="block-content block-content-full text-right border-top">
                    <button type="button" class="btn btn-sm btn-light btn-back-app" @click="abort" 
                        data-dismiss="modal">{{$t("form.all.btn.back")}}</button>
                    <button type="button" class="btn btn-sm btn-primary btn-save-app" @click="save" data-dismiss="modal">
                        <i class="fa fa-check mr-1"></i>{{$t("form.all.btn.save")}}
                    </button>
                </div>
            </div>
        </div>
    </div>
    </div>
    <!-- END Modal Dialog -->
`,
  data() {
    return {
      editFormComponent: "gek-entity-edit-form-" + this.entity,
    };
  },
  methods: {
    abort() {
      this.$emit("entity-edit-abort-" + this.entity);
    },
    save() {
      this.$emit("entity-edit-save-" + this.entity);
    },
  },
  computed: {
    ...Vuex.mapState([
      'entityStores'
    ]),
    editModalId() {
      return this.entity + "EditModal"
    }
  },
});
