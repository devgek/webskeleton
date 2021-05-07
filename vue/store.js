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
      message: null,
    };
  },
  mutations: {
    SET_MESSAGE(state, message) {
      state.message = message;
    },
    SET_USERS(state, users) {
      state.users = users;
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
    setMessage({ commit }, message) {
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
      UserService.getUsers(commit)
    },
    createUser({ dispatch }, userObject) {
      UserService.createUser(dispatch, userObject);
    },
    updateUser({ dispatch }, userObject) {
      UserService.updateUser(dispatch, userObject);
    },
    deleteUser({ dispatch }, userObject) {
      UserService.deleteUser(dispatch, userObject);
    },
    loadContacts({ commit }) {
      ContactService.getContacts(commit)
    },
    createContact({ dispatch }, contactObject) {
      ContactService.createContact(dispatch, contactObject);
    },
    updateContact({ dispatch }, contactObject) {
      ContactService.updateContact(dispatch, contactObject);
    },
    deleteContact({ dispatch }, contactObject) {
      ContactService.deleteContact(dispatch, contactObject);
    },
  },
  getters: {},
});
