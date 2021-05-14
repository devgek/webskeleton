function EntityView(entity, entityName, entityObject, vueInstance) {
    this.entity = entity;
    this.entityName = entityName;
    this.entityObject = entityObject
    this.vue = vueInstance;
    this.editNew = false;
    this.editHeaderNew = entityName + " neu anlegen"
    this.editHeaderUpdate = entityName + " ändern"

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

    this.isEditNew = function() {
      return this.editNew;
    };

    this.editHeaderNew = function(entityName) {
        return entityName + " neu anlegen"
    }

    this.editHeaderUpdate = function(entityName) {
        return entityName + " ändern"
    }
    
};
