// Ready translated locale messages
const messages = {
  de: {
    app: {
      title: "Webskeleton Vue"
    },
    msg: {
      entity: {
        success: {
          create: "{entityDesc} wurde angelegt",
          update: "{entityDesc} wurde geändert",
          delete: "{entityDesc} wurde gelöscht"
        },
        error: {
          create: "{entityDesc} konnte nicht angelegt werden",
          update: "{entityDesc} konnte nicht geändert werden",
          delete: "{entityDesc} konnte nicht gelöscht werden"
        }
      }
    },
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
      login: {
        header: "Bitte anmelden",
        label: {
          user: "Benutzer",
          password: "Passwort",
        },
        button: {
          login: "Anmelden"
        },
        msg: {
          inputrequired: "Username und Passwort müssen angegeben werden."
        }
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
