// Create a new store instance.
const store = new Vuex.Store({
  state() {
    return {
      user: null,
      isAdmin: true,
      users: [],
      contacts: [],
      editUser: {},
      editContact: {}
    };
  },
  mutations: {
    SET_EDIT_USER_ROLE(state, role) {
      state.editUser.Role = role;
    },
    SET_EDIT_USER_EMAIL(state, email) {
      state.editUser.Email = email;
    },
    SET_EDIT_USER_PASS(state, pass) {
      state.editUser.Pass = pass;
    },
    SET_EDIT_USER_NAME(state, name) {
      state.editUser.Name = name;
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
    updateEditUser({commit}, userObject) {
      commit("SET_EDIT_USER", userObject)
    }
  },
  getters: {},
});
