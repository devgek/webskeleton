// Ready translated locale messages
const messages = {
  de: {
    nav: {
      pages: {
        header: "Benutzerfunktionen",
        page1: "Seite 1",
      },
      admin: {
        header: "Administration",
        user: "Benutzer",
        contact: "Kontakt",
      },
    },
    form: {
      all: {
        label: {
          actions: "Aktionen",
        },
        btn: {
          back: "Zurück",
          abort: "Abbrechen",
          save: "Speichern",
          delete: "Löschen",
        },
      },
      page1: {
        header: "Seite 1",
        content: "Das ist der Inhalt von Seite 1",
        labelContacttype: "Kontakttyp:",
      },
      user: {
        list: {
          header: "Benutzer",
          buttonnew: "Neuer Benutzer",
        },
        edit: {
          header: "Benutzer bearbeiten",
          headernew: "Benutzer neu anlegen",
          label: {
            name: "Name:",
            pass: "Passwort:",
            email: "Email:",
            role: "Benutzerrolle:",
            customer: "Anzeige Energiedaten:",
          },
        },
      },
      contact: {
        list: {
          header: "Kontakt",
          buttonnew: "Neuer Kontakt",
          orgtype: "Typ",
          name: "Name",
          nameext: "Namenszusatz",
          contacttype: "Kontakttyp",
          id: "Id",
        },
        edit: {
        header: "Kontakt bearbeiten",
        headernew: "Kontakt neu anlegen",
        label: {
          orgtype: "Organisationstyp:",
          name: "Name:",
          nameext: "Namenszusatz:",
          contacttype: "Kontakttyp:",
          id: "Id:",
        }
      },
    },
  },
}
};

// Create VueI18n instance with options
const i18n = new VueI18n({
  locale: "de", // set locale
  messages, // set locale messages
});
