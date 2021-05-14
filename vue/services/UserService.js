const UserService = {
  getUsers(commit) {
    return axios
      .post("//localhost:8080/api/entitylistuser")
      .then(({ data }) => {
        commit("SET_ENTITY_LIST", {entityName: 'User', entityList: data.EntityObject});
      });
  },
  createUser(dispatch, userObject) {
    return axios
      .post("//localhost:8080/api/entitynewuser", userObject)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Benutzer wurde angelegt",
        };
        dispatch("setMessage", message);
        dispatch("loadEntities", {entityName: 'User'});
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Benutzer wurde nicht angelegt:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
  updateUser(dispatch, userObject) {
    return axios
      .post("//localhost:8080/api/entityedituser", userObject)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Benutzer wurde geändert",
        };
        dispatch("setMessage", message);
        dispatch("loadEntities", {entityName: 'User'});
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Benutzer wurde nicht geändert:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
  deleteUser(dispatch, userObject) {
    return axios
      .post("//localhost:8080/api/entitydeleteuser/" + userObject.ID)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Benutzer wurde gelöscht",
        };
        dispatch("setMessage", message);
        dispatch("loadEntities", {entityName: 'User'});
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Benutzer wurde nicht gelöscht:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
};
