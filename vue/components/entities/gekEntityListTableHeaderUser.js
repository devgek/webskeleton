Vue.component("gek-entity-list-table-header-user", {
  name: "gek-entity-list-table-header-user",
  props: {
  },
  template:
    /*html*/
  `<!-- EntityListTableHeaderUser -->
  <tr>
    <th scope="col">Name</th>
    <th scope="col">Passwort</th>
    <th scope="col">Email</th>
    <th scope="col">Benutzerrolle</th>
    <th scope="col" class="w-5">{{$t("form.all.label.actions")}}</th>
</tr>
<!-- END EntityListTableHeaderUser -->
`,
  data() {
    return {
    };
  },
  methods: {
  },
  computed: {
  },
});
