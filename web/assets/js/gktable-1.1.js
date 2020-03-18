//GKTableObjects container holding all GKTable-objects of current site
var GKTableObjects = [];
//buttonColTemplate code template for creating the edit button column in each table row
var buttonColTemplate = `
<td class="gk-col-buttons">
<div class="text-nowrap">
<button type="button" class="btn btn-secondary gk-btn-edit" onclick="tableId.handleEdit(this);">Bearbeiten</button>
<button type="button" class="btn btn-primary gk-btn-save btn-sm d-none" onclick="tableId.handleSave(this);">Speichern</button>
<button type="button" class="btn btn-primary gk-btn-delete btn-sm d-none" onclick="tableId.handleDelete(this);">Löschen</button>
<button type="button" class="btn btn-primary gk-btn-abort btn-sm d-none" onclick="tableId.handleAbort(this);">Abbrechen</button>
</div> 
</td>`;
//buttonNewTemplate code template for creating the new button
var buttonNewTemplate = `
<button type="button" class="btn btn-secondary gk-btn-new" onclick="tableId.handleNew();">Neu</button>`;

//initialize the GKTable objects after loading site
$(document).ready(function () {
  for (index = 0; index < GKTableObjects.length; ++index) {
    GKTableObjects[index].onLoadDocument();
  }
});

//GKTable constructor function
//each GKTable must have a tableId, a callback function for saving row data and a callback function for deleting row data
function GKTable(tableId, isInlineEditing, saveFunc, deleteFunc) {
  this.tableId = tableId;
  this.isInlineEditing = isInlineEditing
  this.rowData = [];
  this.selectedRow = -1;

  if (saveFunc === undefined) {
    this.saveFunc = this.defaultSaveFunction;
  }
  else {
    this.saveFunc = saveFunc;
  }

  if (deleteFunc === undefined) {
    this.deleteFunc = this.defaultDeleteFunction;
  }
  else {
    this.deleteFunc = deleteFunc;
  }

  this.root = $("table.gk-table[id=" + this.tableId + "]");
  this.buttonCol = buttonColTemplate.replace(
    new RegExp("tableId", "g"),
    this.tableId
  );
  this.buttonNew = buttonNewTemplate.replace(
    new RegExp("tableId", "g"),
    this.tableId
  );
  this.toast = $(".toast[data-tableId=" + this.tableId + "]");

  GKTableObjects.push(this);

  this.selectEnlosingRow = function(obj) {
    this.selectedRow = this.getRowFromParent(obj);
  }

  this.onLoadDocument = function () {
    //prepare gk-table editing
    const $rows = $(this.root).find("tbody tr");
    var theButtonCol = this.buttonCol;
    $rows.each(function (index, element) {
      $(element).addClass("gk-row-edit");
      $(element).attr("data-persisted", "true");

      const $cols = $(element).find("td");
      $cols.each(function (i, c) {
        $(c).addClass("gk-col-edit");
      });

      //add the edit buttons, if isInlineEditing
      if (this.isInlineEditing) {
        $(element).append(theButtonCol);
      }
    });

    //compute the row index
    this.reindexRows();

    //add button new, if isInlineEditing
    if (this.isInlineEditing) {
      $(".gk-btn-new-container[data-tableId=" + this.tableId + "]").prepend(
        this.buttonNew
      );
    }

    //init the toast
    this.toast.toast((delay = 2000));
  };

  this.activateRowEditing = function (theRow) {
    this.disableActions();
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-edit")
      .removeClass("d-inline")
      .addClass("d-none");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-save")
      .removeClass("d-none")
      .addClass("d-inline");
    if (this.isPersisted(theRow)) {
      $(this.root)
        .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-delete")
        .removeClass("d-none")
        .addClass("d-inline");
    }
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-abort")
      .removeClass("d-none")
      .addClass("d-inline");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "]")
      .addClass("bg-warning");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "]")
      .find("td")
      .prop("contenteditable", "true");

    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] td.gk-col-edit")[0]
      .focus();
  };

  this.deactivateRowEditing = function (theRow) {
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-edit")
      .removeClass("d-none")
      .addClass("d-inline");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-new")
      .removeClass("d-none")
      .addClass("d-inline");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-save")
      .removeClass("d-inline")
      .addClass("d-none");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-delete")
      .removeClass("d-inline")
      .addClass("d-none");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "] button.gk-btn-abort")
      .removeClass("d-inline")
      .addClass("d-none");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "]")
      .removeClass("bg-warning");
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "]")
      .find("td")
      .prop("contenteditable", "false");
    this.enableActions();
  };

  this.getRowData = function (theRow) {
    const $cols = $(this.root).find(
      ".gk-row-edit[data-linenr=" + theRow + "] td.gk-col-edit"
    );
    var rowData = [];
    $cols.each(function (index, element) {
      rowData.push($(element).text());
    });

    return rowData;
  };

  this.storeRowData = function (theRow) {
    const $cols = $(this.root).find(
      ".gk-row-edit[data-linenr=" + theRow + "] td.gk-col-edit"
    );
    var theArray = this.rowData;
    $cols.each(function (index, element) {
      theArray.push($(element).text());
    });
  };

  this.restoreRowData = function (theRow) {
    const $cols = $(this.root).find(
      ".gk-row-edit[data-linenr=" + theRow + "] td.gk-col-edit"
    );
    var theArray = this.rowData;
    $cols.each(function (index, element) {
      $(element).text(theArray[index]);
    });
  };

  this.addRow = function (theRow) {
    const $lastRow = $(this.root).find("tbody tr.gk-row-edit:last");
    var clonedRow = $lastRow.clone();
    clonedRow.attr("data-linenr", theRow);
    clonedRow.attr("data-persisted", "false");
    clonedRow.find("td.gk-col-edit").empty();
    $(this.root)
      .find("tbody")
      .append(clonedRow);
  };

  this.enableActions = function () {
    $(this.root)
      .find(".gk-btn-edit")
      .prop("disabled", false);
    $(this.root)
      .find(".gk-btn-new")
      .prop("disabled", false);
  };

  this.disableActions = function () {
    $(this.root)
      .find(".gk-btn-edit")
      .prop("disabled", true);
    $(this.root)
      .find(".gk-btn-new")
      .prop("disabled", true);
  };

  this.deleteRow = function (theRow) {
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "]")
      .remove();
    this.reindexRows();
    this.enableActions();
  };

  this.setPersisted = function (theRow) {
    $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "]")
      .attr("data-persisted", "true");
  };

  this.isPersisted = function (theRow) {
    var persisted = $(this.root)
      .find(".gk-row-edit[data-linenr=" + theRow + "]")
      .attr("data-persisted");
    return persisted == "true";
  };

  this.reindexRows = function () {
    const $rows = $(this.root).find("tbody tr.gk-row-edit");
    $rows.each(function (index, element) {
      $(element).attr("data-linenr", index + 1);
    });
  };

  this.getRowFromParent = function (innerObject) {
    return $(innerObject)
      .parents("tr.gk-row-edit")
      .attr("data-linenr");
  };

  this.handleNew = function () {
    var rows = $(this.root).find(".gk-row-edit").length;
    this.addRow(rows + 1);
    this.activateRowEditing(rows + 1);
  };

  this.handleEdit = function (obj) {
    var line = this.getRowFromParent(obj);
    this.storeRowData(line);
    this.activateRowEditing(line);
  };

  this.handleSave = function (obj) {
    var line = this.getRowFromParent(obj);
    var saved = this.saveFunc(line, this.getRowData(line));
    if (saved) {
      this.setPersisted(line);
      this.deactivateRowEditing(line);
    }
  };

  this.handleDelete = function (obj) {
    var line = this.getRowFromParent(obj);
    var deleted = this.deleteFunc(line);
    if (deleted) {
      this.deleteRow(line);
    }
  };

  this.handleAbort = function (obj) {
    var theRow = this.getRowFromParent(obj);
    if (this.isPersisted(theRow)) {
      this.restoreRowData(theRow);
      this.deactivateRowEditing(theRow);
    } else {
      this.deleteRow(theRow);
    }
  };

  this.showMessage = function (msg) {
    const $toast = $(".toast[data-tableId=" + this.tableId + "]");
    $toast.find("span.gk-toast-text").text(msg);
    $toast.toast("show");
  };

  this.defaultSaveFunction = function(theRow, rowData) {
//do nothing here 
  };
  this.defaultDeleteFunction = function(theRow) {
//do nothing here 
  };
}
