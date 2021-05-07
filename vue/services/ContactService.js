const ContactService = {
  getContacts(commit) {
    return axios
      .post("//localhost:8080/api/entitylistcontact")
      .then(({ data }) => {
        commit("SET_CONTACTS", data.EntityObject);
      });
  },
  createContact(dispatch, contactObject) {
    return axios
      .post("//localhost:8080/api/entitynewcontact", contactObject)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Kontakt wurde angelegt",
        };
        dispatch("setMessage", message);
        dispatch("loadContacts");
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Kontakt wurde nicht angelegt:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
  updateContact(dispatch, contactObject) {
    return axios
      .post("//localhost:8080/api/entityeditcontact", contactObject)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Kontakt wurde geändert",
        };
        dispatch("setMessage", message);
        dispatch("loadContacts");
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Kontakt wurde nicht geändert:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
  deleteContact(dispatch, contactObject) {
    return axios
      .post("//localhost:8080/api/entitydeletecontact/" + contactObject.ID)
      .then(({ data }) => {
        const message = {
          type: "success",
          msg: "Kontakt wurde gelöscht",
        };
        dispatch("setMessage", message);
        dispatch("loadContacts");
      })
      .catch((error) => {
        const message = {
          type: "error",
          msg: "Kontakt wurde nicht gelöscht:" + error.message,
        };
        dispatch("setMessage", message);
      });
  },
};
