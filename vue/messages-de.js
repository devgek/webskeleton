// Ready translated locale messages
const messages = {
  de: {
    nav: {
        pages: {
            header: "Benutzerfunktionen",
            page1: "Seite 1"
        },
        admin: {
            header: "Administration",
            user: "Benutzer",
            contact: "Kontakt"
        }
    },
  },
};

// Create VueI18n instance with options
const i18n = new VueI18n({
  locale: "de", // set locale
  messages, // set locale messages
});
