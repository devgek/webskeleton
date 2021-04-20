// Create a new store instance.
const store = new Vuex.Store({
  state() {
    return {
      user: null,
      isAdmin: true,
      users: [],
      contacts: [],
    };
  },
  mutations: {
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
  },
  getters: {},
});
