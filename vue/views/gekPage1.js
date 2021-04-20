const gekPage1View = Vue.component("gek-page1", {
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
            <h1 class="block-title text-primary">{{$t("form.page1.header")}}</h1>
        </div>
        <div class="block-content block-content-full font-size-sm">
            <form name="formpage1">
                <div class="form-inline">
                    <div class="input-group font-size-sm">
                        <div class="input-group-prepend font-size-sm">
                            <span class="input-group-text font-size-sm bg-grey-light">
                                {{$t("form.page1.labelContacttype")}}
                            </span>
                        </div>
                        <select name="page1FilterContactType" class="form-control font-size-sm bg-info-light"
                            id="page1FilterContactType" @change="onChangeContactType" v-model="selectedContactType">
                            <option v-for="option in getContactTypes" :value="option">{{ option }}</option>
                        </select>
                    </div>
                </div>
            </form>
        </div>
      </div>
      <div class="block block-rounded">
        <div class="block-content block-content-full font-size-sm">
            <div>{{$t("form.page1.content")}}</div>
            <div>{{ selectedContactType }}
        </div>
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
  data() {
    return {
      selectedContactType: "Partner",
    };
  },
  methods: {
    onChangeContactType() {
      console.log("contacttype selected:" + this.selectedContactType);
    },
  },
  computed: {
    getContactTypes() {
      return gkwebapp_T_ContactTypes;
    },
  },
});
