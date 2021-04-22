//create GKTable without inline editing
var userTable = new GKEntityTable("user");

activeGKEntityTable.prepareEditDialog = function () {
  if (activeGKEntityTable.isEditNew()) {
    $("#userEditModalLabel").html("Benutzer neu anlegen");
    $("#userEditName").prop("readonly", false);
    $("#userEditPass").prop("readonly", false);
  } else {
    $("#userEditModalLabel").html("Benutzer ändern");
    $("#userEditName").prop("readonly", true);
    $("#userEditPass").prop("readonly", true);
  }

    store.state.editUser.Name = activeGKEntityTable.editRowData[0];
    store.state.editUser.Pass = activeGKEntityTable.editRowDataHidden[1];
    store.state.editUser.Email = activeGKEntityTable.editRowDataHidden[2];
    store.state.editUser.Role = activeGKEntityTable.editRowDataHidden[3];
};
activeGKEntityTable.prepareSendRowData = function () {
  console.log("getUserView:prepareSendRowData:" + JSON.stringify(store.state.editUser))
  return store.state.editUser
};
activeGKEntityTable.getRowDataFromEntity = function (data) {
  var rowData = [];
  var roleName = gkwebapp_T_RoleTypes[data.EntityObject.Role];
  rowData.push(
    data.EntityObject.Name,
    "********",
    data.EntityObject.Email,
    roleName
  );
  return rowData;
};
activeGKEntityTable.getRowDataHiddenFromEntity = function (data) {
  var rowDataHidden = [];
  rowDataHidden.push(
    data.EntityObject.Name,
    data.EntityObject.Pass,
    data.EntityObject.Email,
    data.EntityObject.Role
  );
  return rowDataHidden;
};