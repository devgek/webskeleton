Vue.component("gek-confirm-delete", {
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
    `<!-- Modal Dialog-->
<div class="modal fade" id="confirmDeleteModal" data-backdrop="static" tabindex="-1" role="dialog"
aria-labelledby="confirmDeleteModalLabel" aria-hidden="true">
<div class="modal-dialog" role="document">
    <div class="modal-content">
        <div class="block block-themed block-transparent mb-0">
            <div class="block-header bg-primary">
                <h3 class="block-title" id="confirmDeleteModalLabel">{{ title }}</h3>
                <div class="block-options">
                  <button type="button" class="btn-block-option" data-dismiss="modal" aria-label="Close">
                      <i class="fa fa-fw fa-times"></i>
                  </button>
                </div>
            </div>
            <div class="block-content font-size-sm">
                <div>{{ confirmationMessage }}</diV>
            </div>

            <div class="block-content block-content-full text-right border-top">
                <button type="button" class="btn btn-sm btn-light btn-back-app"
                    data-dismiss="modal" @click="abort">{{$t("form.all.btn.abort")}}</button>
                <button type="button" class="btn btn-sm btn-primary btn-delete-app" data-dismiss="modal" @click="confirmDelete">
                    <i class="fa fa-check mr-1"></i>{{$t("form.all.btn.delete")}}
                </button>
            </div>
        </div>
    </div>
</div>
</div>
`,
  data() {
    return {
    };
  },
  methods: {
    abort() {
      this.$emit("entity-delete-abort-" + this.entity);
    },
    confirmDelete() {
      this.$emit("entity-delete-confirm-" + this.entity);
    },
  },
  computed: {
    title() {
      return this.entityName + " löschen";
    },
    confirmationMessage() {
      return this.entityName + " wirklich löschen?";
    },
  },
});
