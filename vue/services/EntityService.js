const EntityService = {
  getEntities(commit, payload) {
    return axios
      .post("//localhost:8080/api/entitylist" + payload.entityName)
      .then(({ data }) => {
        commit("SET_ENTITY_LIST", {entityName: payload.entityName, entityList: data.EntityObject});
      });
  },
  createEntity(dispatch, payload) {
    return axios
      .post("//localhost:8080/api/entitynew" + payload.entityName, payload.entityObject)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Entität wurde angelegt",
        };
        dispatch("setMessage", message);
        dispatch("loadEntities", {entityName: payload.entityName});
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Entität wurde nicht angelegt:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
  updateEntity(dispatch, payload) {
    return axios
      .post("//localhost:8080/api/entityedit" + payload.entityName, payload.entityObject)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Entität wurde geändert",
        };
        dispatch("setMessage", message);
        dispatch("loadEntities", {entityName: payload.entityName});
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Entität wurde nicht geändert:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
  deleteEntity(dispatch, payload) {
    return axios
      .post("//localhost:8080/api/entitydelete" + payload.entityName + "/" + payload.entityObject.ID)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Entität wurde gelöscht",
        };
        dispatch("setMessage", message);
        dispatch("loadEntities", {entityName: payload.entityName});
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Entität wurde nicht gelöscht:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
};
