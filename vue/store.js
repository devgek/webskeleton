// Create a new store instance.
const store = new Vuex.Store({
  state() {
    return {
      user: null,
      isAdmin: true,
      users: [],
      contacts: [],
      editUser: {},
      editContact: {},
      message: null
    };
  },
  mutations: {
    SET_MESSAGE(state, message) {
      state.message = message;
    },
    SET_EDIT_USER(state, user) {
      state.editUser = user;
    },
    SET_USERS(state, users) {
      state.users = users;
    },
    SET_EDIT_CONTACT(state, contact) {
      state.editContact = contact;
    },
    SET_CONTACTS(state, contacts) {
      state.contacts = contacts;
    },
    SET_USER_DATA(state, userData) {
      localStorage.setItem("user", JSON.stringify(userData));
      axios.defaults.headers.common[
        "Authorization"
      ] = `Bearer ${userData.token}`;
      state.user = userData;
    },
    SET_ADMIN(state, isAdmin) {
      state.isAdmin = isAdmin;
    },
    LOGOUT() {
      localStorage.removeItem("user");
      location.reload();
    },
  },
  actions: {
    setMessage({commit}, message) {
      commit("SET_MESSAGE", message);
    },
    login({ commit }, credentials) {
      return axios
        .post("//localhost:8080/api/login", credentials)
        .then(({ data }) => {
          commit("SET_USER_DATA", data);
        });
    },
    logout({ commit }) {
      commit("LOGOUT");
    },
    loadUsers({ commit }) {
      return axios
        .post("//localhost:8080/api/entitylistuser")
        .then(({ data }) => {
          commit("SET_USERS", data.EntityObject);
        });
    },
    loadContacts({ commit }) {
      return axios
        .post("//localhost:8080/api/entitylistcontact")
        .then(({ data }) => {
          commit("SET_CONTACTS", data.EntityObject);
        });
    },
    createUser({commit, dispatch}, userObject) {
      return axios
        .post("//localhost:8080/api/entitynewuser", userObject)
        .then(({ data }) => {
          commit("SET_EDIT_USER", data.EntityObject);
          const message = {
            type: "success",
            msg: "Benutzer wurde angelegt"
          }
          dispatch("setMessage", message);
          dispatch("loadUsers");
        })
        .catch((error) => {
          const message = {
            type: "error",
            msg: "Benutzer wurde nicht angelegt:" + error.message
          }
          dispatch("setMessage", message);
        });
      },
      updateUser({commit, dispatch}, userObject) {
        return axios
          .post("//localhost:8080/api/entityedituser", userObject)
          .then(({ data }) => {
            commit("SET_EDIT_USER", data.EntityObject);
            const message = {
              type: "success",
              msg: "Benutzer wurde geändert"
            }
            dispatch("setMessage", message);
            dispatch("loadUsers");
          })
          .catch((error) => {
            const message = {
              type: "error",
              msg: "Benutzer wurde nicht geändert:" + error.message
            }
            dispatch("setMessage", message);
          });
        },
        deleteUser({commit, dispatch}, userObject) {
          return axios
            .post("//localhost:8080/api/entitydeleteuser/" + userObject.id)
            .then(({ data }) => {
              const message = {
                type: "success",
                msg: "Benutzer wurde gelöscht"
              }
              dispatch("setMessage", message);
              dispatch("loadUsers");
            })
            .catch((error) => {
              const message = {
                type: "error",
                msg: "Benutzer wurde nicht gelöscht:" + error.message
              }
              dispatch("setMessage", message);
            });
          },      },
  getters: {}
});
