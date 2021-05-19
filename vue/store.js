// Create a new store instance.
const store = new Vuex.Store({
  state() {
    return {
      user: null,
      isAdmin: true,
      message: null,
      entityStores: {"User": null, "Contact": null}
    };
  },
  mutations: {
    SET_MESSAGE(state, message) {
      state.message = message;
    },
    NEW_ENTITY_STORE(state, payload) {
      state.entityStores[payload.entityName] = new EntityStore(payload.entityName, payload.newEntityObjectFn, this);
    },
    SET_ENTITY_NEW(state, payload) {
      state.entityStores[payload.entityName].entityObject = state.entityStores[payload.entityName].newEntityObjectFn();
      state.entityStores[payload.entityName].editNew = true;
    },
    SET_ENTITY_EDIT(state, payload) {
      entityObjectRef = state.entityStores[payload.entityName].entityList[payload.entityIndex]
      state.entityStores[payload.entityName].entityObject = JSON.parse(JSON.stringify(entityObjectRef));
      state.entityStores[payload.entityName].editNew = false;
    },
    SET_ENTITY_LIST(state, payload) {
      state.entityStores[payload.entityName].entityList = payload.entityList;
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
    LOGOUT(state) {
      localStorage.removeItem("user");
      // location.reload();
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
    startEntityStore({ commit, dispatch }, payload) {
      commit("NEW_ENTITY_STORE", payload)

      dispatch("loadEntities", payload)
    },
    loadEntities({ commit }, payload) {
      EntityService.getEntities(commit, payload);
    },
    saveEntity({ dispatch, getters }, payload) {
      payload.entityObject = getters.getEditEntityObjectByEntityName(payload.entityName);
      if (getters.getEditNewByEntityName(payload.entityName)) {
        EntityService.createEntity(dispatch, payload);
      }
      else {
        EntityService.updateEntity(dispatch, payload);
      }
    },
    deleteEntity({ dispatch, getters }, payload) {
      payload.entityObject = getters.getEditEntityObjectByEntityName(payload.entityName);
      EntityService.deleteEntity(dispatch, payload);
    }
  },
  getters: {
    isAdminUser(state) {
      return state.isAdmin;
    },
    getEntityListByEntityName: (state) => (entityName) => {
      return state.entityStores[entityName].entityList;
    },
    getEditEntityObjectByEntityName: (state) => (entityName) => {
      return state.entityStores[entityName].entityObject;
    },
    getEditNewByEntityName: (state) => (entityName) => {
      return state.entityStores[entityName].editNew;
    },
    getUser(state) {
      return state.user;
    }
  },
});
