Vue.component("gek-entity-list-table-header-contact", {
  name: "gek-entity-list-table-header-contact",
  props: {
  },
  template:
    /*html*/
  `<!-- EntityListTableHeaderContact -->
  <tr>
    <th scope="col">{{$t("form.contact.list.orgtype")}}</th>
    <th scope="col">{{$t("form.contact.list.name")}}</th>
    <th scope="col">{{$t("form.contact.list.nameext")}}</th>
    <th scope="col">{{$t("form.contact.list.contacttype")}}</th>
    <th scope="col">{{$t("form.contact.list.id")}}</th>
    <th scope="col" class="w-5">{{$t("form.all.label.actions")}}</th>
</tr>
<!-- END EntityListTableHeaderContact -->
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
