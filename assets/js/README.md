## gktable-1.1.js Javascript-lib zum Inline-Editieren einer Tabelle

Benötigt: bootstrap-4.3.1, jquery-3.3.1

Zur Verwendung in einer Seite muss:

1.) Das Table-Element muss die Style-Klasse "gk-table" und eine eindeutige tableId haben, Bsp.:

```html
<table class="gk-table" id="inoutTable"></table>
```

2.) Die Konstruktor-Funktion "GKTable" muss mit den Parametern tableId, saveFunc und deleteFunc aufgerufen werden, Bsp:

```javascript
var inoutTable = new GKTable(
  "inoutTable",
  inoutSaveRowRequest,
  inoutDeleteRowRequest
);
```

Durch eine eindeutige tableId ist es möglich, dass auch mehr als eine GKTable auf einer Seite vorhanden ist.

saveFunc ist eine Callbackfunktion, die mit der aktuellen Zeilennummer und den Daten aufgerufen wird, damit die Daten auch zum Server geposted werden können. Die Funktion muss true/false (erfolgreich gespeichert) zurückliefern.
Bsp.:

```javascript
function inoutSaveRowRequest(theRow, rowData) {
  var tdArt = rowData[0];
  var tdText = rowData[1];
  var tdKonto = rowData[2];
  var tdGkonto = rowData[3];
  var tdSteuer = rowData[4];

  var url = "/saveinout";
  var result = false;

  var posting = $.post(url, {
    inoutLine: theRow,
    inoutArt: tdArt,
    inoutText: tdText,
    inoutKonto: tdKonto,
    inoutGkonto: tdGkonto,
    inoutSteuer: tdSteuer
  });

  posting.done(function(data) {
    var isError = data.Errorflag;
    if (isError) {
      inoutTable.showMessage(data.Message);
    } else {
      inoutTable.showMessage("Zuordnung wurde gespeichert");
      result = true;
    }
  });

  posting.fail(function(xhr, status, error) {
    inoutTable.showMessage(
      "Technischer Fehler: " + xhr.status + ":" + xhr.statusText
    );
  });

  return result;
}
```

deleteFunc ist eine Callbackfunktion, die mit der aktuellen Zeilennummer aufgerufen wird, damit die Daten auch zum Server geposted werden können.
Die Funktion muss true/false (erfolgreich gelöscht) zurückliefern.
Bsp.:

```javascript
function inoutDeleteRowRequest(theRow) {
  var url = "/deleteinout";
  var result = false;

  var posting = $.post(url, { inoutLine: theRow });

  posting.done(function(data) {
    var isError = data.Errorflag;
    if (isError) {
      inoutTable.showMessage(data.Message);
    } else {
      inoutTable.showMessage("Zuordnung wurde gelöscht");
      result = true;
    }
  });

  posting.fail(function(xhr, status, error) {
    inoutTable.showMessage(
      "Technischer Fehler: " + xhr.status + ":" + xhr.statusText
    );
  });

  return result;
}
```
