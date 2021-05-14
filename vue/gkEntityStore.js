function EntityStore(entityName, newEntityObjectFn, vuex) {
    var entityStore = this;
    this.entityName = entityName;
    this.newEntityObjectFn = newEntityObjectFn;
    this.entityObject = this.newEntityObjectFn.call();
    this.entityList = null;
    this.vuex = vuex;
    this.editNew = false;

    this.doSave = function () {
      console.log("doSave", this)
      if (this.editNew) {
        this.vue.$store.dispatch("create" + this.entityName, this.entityObject);
      } else {
        this.vue.$store.dispatch("update" + this.entityName, this.entityObject);
      }
    };

    this.doDelete = function (confirmed) {
      if (confirmed) {
        this.vue.$store.dispatch("delete" + this.entityName, this.entityObject);
      }
    };

    var getEditHeader = entityStore.getEditHeader;
    this.getEditHeader = function boundGetEditHeader() {
        return getEditHeader.call(entityStore)
    }
    
};

EntityStore.prototype.getEditHeader = function getEditHeader() {
  if (this.editNew) {
    return this.entityName + " neu anlegen"
  }
  else {
    return this.entityName + " ändern"
  }
}