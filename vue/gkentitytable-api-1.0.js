//GKEntityTable with REST-API 1.0 (for vue.js-Frontend)
//GKEntityTableObjects container holding all GKTable-objects of current site
var GKEntityTableObjects = [];
var activeGKEntityTable;

//initialize the GKTable objects after loading site
// $(document).ready(function () {
//   for (index = 0; index < GKEntityTableObjects.length; ++index) {
//     GKEntityTableObjects[index].onLoadDocument();
//   }
// });

//GKEntityTable constructor function
//each GKEntityTable must have a tableId, an entity and member functions(getRowDataFromEntity, getRowDataHiddenFromEntity) for getting rowData and rowDataHidden from entity
function GKEntityTable(entity, entityEmbedded) {
  this.initialized = false;
  this.entity = entity;
  this.tableId = entity + "Table";
  this.dialogId = entity + "EditModal";
  this.urlEdit = "/api/entityedit" + this.entity;
  this.urlNew = "/api/entitynew" + this.entity;
  this.urlDelete = "/api/entitydelete" + this.entity;
  this.editRowData = [];
  this.editRowDataHidden = [];
  this.editRow = {};
  this.editRowKey = 0;
  this.editNew = false;

  if (entityEmbedded == undefined) {
    this.entityEmbedded = "";
    this.dialogIdEmbedded = "";
  } else {
    this.entityEmbedded = entityEmbedded;
    this.dialogIdEmbedded = entityEmbedded + "EditModal";
  }
  this.urlEditEmbedded = "entityedit" + this.entityEmbedded;
  this.urlNewEmbedded = "entitynew" + this.entityEmbedded;
  this.urlDeleteEmbedded = "entitydelete" + this.entityEmbedded;
  this.editRowDataEmbedded = [];
  this.editRowDataHiddenEmbedded = [];
  this.editRowEmbedded = {};
  this.editRowKeyEmbedded = 0;
  this.editNewEmbedded = false;

  this.entityOptionsArray = [];
  this.entityOptionsName = "";
  this.colFilterValues = [];
  this.colFilterColumns = [];

  this.root = null;

  GKEntityTableObjects.push(this);
  activeGKEntityTable = this;

  this.initialize = function () {
    console.log("initialize GKEntityTable " + this.tableId)

    this.root = $("table.gk-table[id=" + this.tableId + "]");

    //prepare gk-table editing
    const $rows = $(this.root).find(
      "tbody tr:not(.gk-row-section)[data-entityid]"
    );
    if ($rows.length < 1) {
      // table rows not loaded yet, so return
      console.log("table rows not loaded yet")
      return;
    }

    if (this.initialized) {
      console.log("already initialized");
      return;
    }

    $rows.each(function (index, element) {
      $(element).addClass("gk-row-edit");

      const $cols = $(element).find("td[data-gkvval]");
      $cols.each(function (i, c) {
        $(c).addClass("gk-col-edit gk-col-" + i);
      });
    });

    const $rowsEmbedded = $(this.root).find(
      "tbody tr.gk-row-section[data-entityidembedded]"
    );
    $rowsEmbedded.each(function (index, element) {
      $(element).addClass("gk-row-embedded");

      const $cols = $(element).find("td[data-gkvval]");
      $cols.each(function (i, c) {
        $(c).addClass("gk-col-edit gk-col-" + i);
      });
    });

    //prepare edit buttons
    const $editButtons = $(
      "button.gk-btn-edit, button.gk-btn-new, button.gk-btn-delete"
    );
    /*
    $editButtons.each(function (index, element) {
      $(element).click(function () {
        activeGKEntityTable.onStartRowEditing(this);
      });
    });
    */
    //prepare modal dialog for row editing
    $("#" + this.dialogId).on("shown.bs.modal", function (event) {
      console.log("on shown.bs.modal before prepareEditDialog")
      activeGKEntityTable.prepareEditDialog();
    });

    //prepare modal dialog for deleting
    /*
    $("#confirmDeleteModal .btn-delete-app").click(function () {
      activeGKEntityTable.deleteRowData();
    });
    */
   
    if (this.dialogIdEmbedded != "") {
      //prepare new embedded entities buttons
      const $newEmbeddedButtons = $("button.gk-btn-new-embedded");
      $newEmbeddedButtons.each(function (index, element) {
        $(element).click(function () {
          activeGKEntityTable.onStartRowNewEmbedded(this);
        });
      });

      //prepare delete embedded entities buttons
      const $deleteEmbeddedButtons = $("button.gk-btn-delete-embedded");
      $deleteEmbeddedButtons.each(function (index, element) {
        $(element).click(function () {
          activeGKEntityTable.onStartRowDeleteEmbedded(this);
        });
      });

      //prepare modal dialog for embedded row editing
      $("#" + this.dialogIdEmbedded).on("shown.bs.modal", function (event) {
        activeGKEntityTable.prepareEmbeddedDialog();
      });

      $("#" + this.dialogIdEmbedded + " .btn-save-app").click(function () {
        var sendParams = activeGKEntityTable.prepareSendRowDataEmbedded();
        activeGKEntityTable.sendRowDataEmbedded(sendParams);
      });

      //prepare modal dialog for embedded deleting
      $("#confirmDeleteEmbeddedModal .btn-delete-app").click(function () {
        activeGKEntityTable.deleteRowDataEmbedded();
      });
    }

    this.initialized = true;
  };

  this.getRowData = function (entityId) {
    const $cols = $(this.root).find(
      ".gk-row-edit[data-entityid=" + entityId + "] td.gk-col-edit"
    );
    var rowData = [];
    $cols.each(function (index, element) {
      rowData.push($(element).text());
    });

    return rowData;
  };

  this.getRowDataHidden = function (entityId) {
    const $cols = $(this.root).find(
      ".gk-row-edit[data-entityid=" + entityId + "] td.gk-col-edit"
    );
    var rowData = [];
    $cols.each(function (index, element) {
      rowData.push($(element).attr("data-gkvval"));
    });

    return rowData;
  };

  this.addRow = function (entityId) {
    const $lastRow = $(this.root).find("tbody tr.gk-row-edit:last");
    var clonedRow = $lastRow.clone(true);
    clonedRow.attr("data-entityid", entityId);
    clonedRow.find("td.gk-col-edit:not(:last)").attr("data-gkvval", "").empty();
    $(this.root).find("tbody").append(clonedRow);
  };

  this.addRowEmbedded = function (entityId) {
    var entityRow = this.getEditRow();
    const $nextTbodyInnerTbody = $(entityRow)
      .parent("tbody")
      .next("tbody")
      .find("table tbody");
    const $lastEmbeddedRow = $nextTbodyInnerTbody.find("tr:last");
    var clonedRow = $lastEmbeddedRow.clone(true);
    clonedRow.attr("data-entityidembedded", entityId);
    clonedRow.find("td.gk-col-edit:not(:last)").attr("data-gkvval", "").empty();
    $nextTbodyInnerTbody.append(clonedRow);
  };

  this.deleteEditRow = function () {
    this.editRow.remove();
  };

  this.deleteEditRowEmbedded = function () {
    this.editRowEmbedded.remove();
  };

  this.selectRowForEditing = function (entityId) {
    this.editRow = $(this.root).find(
      ".gk-row-edit[data-entityid=" + entityId + "]"
    );
  };

  this.selectRowForEditingEmbedded = function (entityId) {
    this.editRowEmbedded = $(this.root).find(
      "tr[data-entityidembedded=" + entityId + "]"
    );
  };

  this.changeRowData = function (entityId, data) {
    const $cols = $(this.root).find(
      ".gk-row-edit[data-entityid=" + entityId + "] td.gk-col-edit"
    );
    var theArray = data;
    $cols.each(function (index, element) {
      $(element).text(theArray[index]);
    });
  };

  this.changeRowDataEmbedded = function (entityId, data) {
    const $cols = $(this.root).find(
      "tr.gk-row-section[data-entityidembedded=" + entityId + "] td.gk-col-edit"
    );
    var theArray = data;
    $cols.each(function (index, element) {
      $(element).text(theArray[index]);
    });
  };

  this.changeRowDataHidden = function (entityId, data) {
    const $cols = $(this.root).find(
      ".gk-row-edit[data-entityid=" + entityId + "] td.gk-col-edit"
    );
    var theArray = data;
    $cols.each(function (index, element) {
      $(element).attr("data-gkvval", theArray[index]);
    });
  };

  this.changeRowDataHiddenEmbedded = function (entityId, data) {
    const $cols = $(this.root).find(
      "tr.gk-row-section[data-entityidembedded=" + entityId + "] td.gk-col-edit"
    );
    var theArray = data;
    $cols.each(function (index, element) {
      $(element).attr("data-gkvval", theArray[index]);
    });
  };

  this.getEditRow = function () {
    return this.editRow;
  };

  this.getDeleteUrl = function () {
    return this.urlDelete;
  };

  this.getDeleteEmbeddedUrl = function () {
    return this.urlDeleteEmbedded;
  };

  this.getEditRowKey = function () {
    return this.editRowKey;
  };

  this.getEditRowKeyEmbedded = function () {
    return this.editRowKeyEmbedded;
  };

  this.setEditRowKey = function (rowKey) {
    this.editRowKey = rowKey;
  };

  this.setEditRowKeyEmbedded = function (rowKey) {
    this.editRowKeyEmbedded = rowKey;
  };

  this.isEditNew = function () {
    return this.editNew;
  };

  this.setEditNew = function (isEditNew) {
    this.editNew = isEditNew;
  };

  this.getRowKeyFromParent = function (innerObject) {
    return $(innerObject).parents("tr.gk-row-edit").attr("data-entityid");
  };

  this.getRowKeyEmbeddedFromParent = function (innerObject) {
    return $(innerObject).parents("tr").attr("data-entityidembedded");
  };

  this.setEditRowKeyFromParent = function (innerObject) {
    this.setEditRowKey(this.getRowKeyFromParent(innerObject));
  };

  this.handleNew = function (entityId) {
    this.addRow(entityId);
    this.selectRowForEditing(entityId);
  };

  this.handleNewEmbedded = function (entityId) {
    this.addRowEmbedded(entityId);
  };

  this.loadEntityOptionsForSelect = function (entityName) {
    var posting = $.post("apioptionlist" + entityName);

    posting.done(function (data) {
      var isError = data.IsError;
      if (isError) {
        console.log(
          "Error while loading entity list of ",
          entityName,
          ": ",
          data.Message
        );
      } else {
        activeGKEntityTable.entityOptionsArray[entityName] = data.EntityOptions;
      }
    });

    posting.fail(function (xhr, status, error) {
      console.log("Error while ajax post:" + xhr.status + ":" + xhr.statusText);
    });
  };

  this.prepareEntitySelect = function (
    entityName,
    selectId,
    selectedValue,
    defaultOption
  ) {
    return this.prepareEntitySelectFiltered(
      "$$all$$",
      entityName,
      selectId,
      selectedValue,
      defaultOption
    );
  };

  this.prepareEntitySelectFiltered = function (
    filter,
    entityName,
    selectId,
    selectedValue,
    defaultOption
  ) {
    if (typeof this.entityOptionsArray[entityName] == "undefined") {
      this.loadEntityOptionsForSelect(entityName);
    }

    var theSelect = $("#" + selectId);
    var theOptions = theSelect.prop("options");
    $("option", theSelect).remove();
    if (defaultOption) {
      theOptions[theOptions.length] = defaultOption;
    }

    setTimeout(function () {
      if (
        typeof activeGKEntityTable.entityOptionsArray[entityName] != "undefined"
      ) {
        $.each(activeGKEntityTable.entityOptionsArray[entityName], function (
          index,
          option
        ) {
          if (filter == "$$all$$" || filter == option.Filter) {
            theOptions[theOptions.length] = new Option(option.Value, option.ID);
          }
        });
        theSelect.val(selectedValue);
      } else {
        console.log("no options to load");
      }
    }, 300);
  };

  this.prepareSelect = function (filter, theOptions) {
    if (activeGKEntityTable.entityOptions.length) {
      $.each(activeGKEntityTable.entityOptions, function (index, option) {
        if (filter == "$$all$$" || filter == option.Filter) {
          theOptions[theOptions.length] = new Option(option.Value, option.ID);
        }
      });
    } else {
      console.log("no options to load");
    }
  };

  this.initColFilter = function (filterIdx, colIdx, colVal) {
    this.colFilterValues[filterIdx] = colVal;
    this.colFilterColumns[filterIdx] = colIdx;
  };

  this.applyColFilter = function (filterIdx, colVal) {
    var theTable = this;
    this.colFilterValues[filterIdx] = colVal;

    const $rows = $(this.root).find(
      "tbody tr.gk-row-edit:not(.gk-row-section)"
    );
    $rows.each(function (index, row) {
      const $cols = $(row).find("td.gk-col-edit");

      var activate = true;
      for (i = 0; i < theTable.colFilterValues.length; i++) {
        var filterVal = theTable.colFilterValues[i];
        var colIdx = theTable.colFilterColumns[i];
        var colVal = $($cols[colIdx]).attr("data-gkvval");
        if (filterVal != "0" && filterVal != colVal) {
          activate = false;
        }
      }

      var tBody = $(row).parents("tbody");
      if (activate) {
        $(row).parents("tbody").removeClass("d-none");
      } else {
        if (
          $(tBody).hasClass("js-table-sections-header") &&
          $(tBody).hasClass("show") &&
          $(tBody).hasClass("table-active")
        ) {
          $(tBody).removeClass("show").removeClass("table-active");
        }
        $(tBody).addClass("d-none");
      }
    });
  };

  this.onStartRowEditing = function (trigger) {
    console.log("onStartRowEditing")
    //trigger = button, that started row editing
    var rowKey = this.getRowKeyFromParent(trigger); //get the rowKey -> entityId
    this.setEditRowKey(rowKey);
    var isEditNew = $(trigger).hasClass("gk-btn-new"); //get edit mode (new/change)
    this.setEditNew(isEditNew);
    if (isEditNew) {
      this.editRowData = [];
      this.editRowDataHidden = [];
    } else {
      this.selectRowForEditing(rowKey);
      this.editRowData = activeGKEntityTable.getRowData(rowKey);
      this.editRowDataHidden = activeGKEntityTable.getRowDataHidden(rowKey);
    }
  };

  this.onStartRowNewEmbedded = function (trigger) {
    //trigger = button, that started row editing
    var rowKey = this.getRowKeyFromParent(trigger); //get the embedded rowKey -> entityIdEmbedded
    this.setEditRowKey(rowKey);
    var isEditNew = true; //get edit mode (new/change)
    this.setEditNew(isEditNew);
    this.selectRowForEditing(rowKey);
    this.editRowData = this.getRowData(rowKey);
    this.editRowDataHidden = this.getRowDataHidden(rowKey);
  };

  this.onStartRowDeleteEmbedded = function (trigger) {
    //trigger = button, that started row editing
    var rowKey = this.getRowKeyEmbeddedFromParent(trigger); //get the embedded rowKey -> entityIdEmbedded
    this.setEditRowKeyEmbedded(rowKey);
    this.selectRowForEditingEmbedded(rowKey);
  };

  this.getEditUrl = function () {
    if (this.isEditNew()) {
      return this.urlNew;
    } else {
      return this.urlEdit;
    }
  };

  this.getEditEmbeddedUrl = function () {
    if (this.isEditNew()) {
      return this.urlNewEmbedded;
    } else {
      return this.urlEditEmbedded;
    }
  };

  this.showMessage = function (msg) {
    const $toast = $(".toast[data-tableId='" + this.tableId + "']");
    $toast.find("span.gk-toast-text").text(msg);
    $toast.toast("show");
  };

  this.showConfirmDeleteMessage = function (msg) {
    const $toast = $(".toast[data-toastid='confirmDeleteToast']");
    $toast.find("span.gk-toast-text").text(msg);
    $toast.toast("show");
  };

  this.showConfirmDeleteEmbeddedMessage = function (msg) {
    const $toast = $(".toast[data-toastid='confirmDeleteEmbeddedToast']");
    $toast.find("span.gk-toast-text").text(msg);
    $toast.toast("show");
  };

  this.sendRowData = function (sendParams) {
    /*
    var postData = {};
    for (var key in sendParams) {
      postData[key] = sendParams[key];
    }
    */
   var postData = sendParams;
   console.log("sendRowData:" + JSON.stringify(postData))
    var axiosPost = axios.post(this.getEditUrl(), postData);

    axiosPost.then(function (response) {
      var isError = false;
      if (isError) {
        activeGKEntityTable.showMessage(data.Message);
      } else {
        var entityId = response.data.ID;
        if (activeGKEntityTable.isEditNew()) {
          activeGKEntityTable.handleNew(entityId);
        }
        var rowData = activeGKEntityTable.getRowDataFromEntity(response.data);
        var rowDataHidden = activeGKEntityTable.getRowDataHiddenFromEntity(
          data
        );
        activeGKEntityTable.changeRowData(entityId, rowData);
        if (typeof activeGKEntityTable.onChangeRowData === "function") {
          activeGKEntityTable.onChangeRowData(entityId, rowData);
        }
        activeGKEntityTable.changeRowDataHidden(entityId, rowDataHidden);
      }
    });

    axiosPost.catch(function (err) {
      if (err.response) {
        // The request was made and the server responded with a status code
        // that falls out of the range of 2xx
        activeGKEntityTable.showMessage(err.response.data  + "," + err.response.status + "," + err.response.headers);
      } else if (err.request) {
        // The request was made but no response was received
        // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
        // http.ClientRequest in node.js
        activeGKEntityTable.showMessage(err.request);
      } else {
        // Something happened in setting up the request that triggered an Error
        activeGKEntityTable.showMessage(err.message);
      }
    });
  };

  this.sendRowDataEmbedded = function (sendParams) {
    var postData = {};
    for (var key in sendParams) {
      postData[key] = sendParams[key];
    }

    var posting = $.post(this.getEditEmbeddedUrl(), postData);

    posting.done(function (data) {
      var isError = data.IsError;
      if (isError) {
        activeGKEntityTable.showMessage(data.Message);
      } else {
        var entityId = data.EntityObject.ID;
        if (activeGKEntityTable.isEditNew()) {
          activeGKEntityTable.handleNewEmbedded(entityId);
        }
        var rowData = activeGKEntityTable.getRowDataFromEntityEmbedded(data);
        var rowDataHidden = activeGKEntityTable.getRowDataHiddenFromEntityEmbedded(
          data
        );

        activeGKEntityTable.changeRowDataEmbedded(entityId, rowData);
        if (typeof activeGKEntityTable.onChangeRowDataEmbedded === "function") {
          activeGKEntityTable.onChangeRowDataEmbedded(entityId, rowData);
        }
        activeGKEntityTable.changeRowDataHiddenEmbedded(
          entityId,
          rowDataHidden
        );
        activeGKEntityTable.showMessage(data.Message);
      }
    });

    posting.fail(function (xhr, status, error) {
      activeGKEntityTable.showMessage(
        '{{.Messages.GetString "msg.error.ajax"}}' +
          xhr.status +
          ":" +
          xhr.statusText
      );
    });
  };

  this.deleteRowData = function () {
    var posting = $.post(activeGKEntityTable.getDeleteUrl(), {
      gkvObjId: activeGKEntityTable.getEditRowKey(),
    });

    posting.done(function (data) {
      var isError = data.IsError;
      if (isError) {
        activeGKEntityTable.showConfirmDeleteMessage(data.Message);
      } else {
        var theRow = activeGKEntityTable.getEditRow();
        activeGKEntityTable.deleteEditRow();
        activeGKEntityTable.showConfirmDeleteMessage(data.Message);
      }
    });

    posting.fail(function (xhr, status, error) {
      activeGKEntityTable.showConfirmDeleteMessage(
        "Technischer Fehler: " + xhr.status + ":" + xhr.statusText
      );
    });
  };

  this.deleteRowDataEmbedded = function () {
    var posting = $.post(activeGKEntityTable.getDeleteEmbeddedUrl(), {
      gkvObjId: activeGKEntityTable.getEditRowKeyEmbedded(),
    });

    posting.done(function (data) {
      var isError = data.IsError;
      if (isError) {
        activeGKEntityTable.showConfirmDeleteEmbeddedMessage(data.Message);
      } else {
        activeGKEntityTable.deleteEditRowEmbedded();
        activeGKEntityTable.showConfirmDeleteEmbeddedMessage(data.Message);
      }
    });

    posting.fail(function (xhr, status, error) {
      activeGKEntityTable.showConfirmDeleteEmbeddedMessage(
        "Technischer Fehler: " + xhr.status + ":" + xhr.statusText
      );
    });
  };
} //end of GKEntityTable
