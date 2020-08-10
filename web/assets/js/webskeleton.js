//bepo-types
//converting type values (int) to string representation of the type
//used in javascript code on client side
var bepo_T_OrgTypes = ["ORG", "PER"];
var bepo_T_CustomerTypes = [
  "Kunde",
  "Lieferant",
  "Partner",
  "Interessent",
  "Werbung",
];
var bepo_T_MenuTypes = ["admin", "consumption"];
var bepo_T_EnergyTypes = [
  "Undefiniert",
  "Strom",
  "Nutzwasser",
  "Druckluft",
  "Trinkwasser",
  "Erdgas",
  "Diesel",
  "Heizung",
];
var bepo_T_MeterTypes = ["Stromzähler", "Anderer Zähler"];
var bepo_T_RoleTypes = ["Kunde", "BAGBenutzer", "BAGAdministrator"];

function bepo_prepareSelect(selectId, selectedValue) {
  var theSelect = $("#" + selectId);

  theSelect.val(selectedValue);
}

function bepo_modalShowMessage(modalId, msg) {
  const $toast = $("#" + modalId + " .toast");
  $toast.find("span.gk-toast-text").text(msg);
  $toast.toast("show");
}

function bepo_format_curr(num) {
  var str = num.toString().replace("$", ""),
    parts = false,
    output = [],
    i = 1,
    formatted = null;
  if (str.indexOf(".") > 0) {
    parts = str.split(".");
    str = parts[0];
  }
  str = str.trim().split("").reverse();
  for (var j = 0, len = str.length; j < len; j++) {
    if (str[j] != ".") {
      output.push(str[j]);
      if (i % 3 == 0 && j < len - 1) {
        output.push(".");
      }
      i++;
    }
  }
  formatted = output.reverse().join("");
  if (formatted.substr(0, 2) == "-.") {
    formatted = "-" + formatted.substr(2);
  }
  return formatted + (parts ? "," + parts[1].substr(0, 2) : "");
}

function bepo_format_all_curr() {
  const $cols = $(".gk-format-curr");
  $cols.each(function (index, element) {
    var vf = bepo_format_curr($(element).text());
    $(element).text(vf);
  });
}
