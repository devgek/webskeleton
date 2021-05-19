// Routes
const routes = [
  {
    path: "/login",
    alias: "",
    component: gekLoginView,
    name: "Login",
    props: { mainHeader: "Bitte anmelden", startPage: "Page1" },
  },
  {
    path: "/",
    component: gekLayoutView,
    children: [
      {
        path: "home",
        component: gekHomeView,
        name: "Home",
        meta: { description: "show home" },
      },
      {
        path: "page1",
        component: gekPage1View,
        name: "Page1",
        meta: { description: "show page1" },
      },
    ],
  },
  {
    path: "/admin",
    component: gekLayoutView,
    children: [
      {
        path: "user",
        component: gekUserView,
        name: "User",
        meta: { description: "show user" },
      },
      {
        path: "contact",
        component: gekContactView,
        name: "Contact",
        meta: { description: "show contact" },
      },
    ],
  },
  {
    // not found handler
    path: "*",
    component: gekNotFoundView,
  },
];

const router = new VueRouter({
  routes,
});

router.beforeEach((to, from, next) => {
  // redirect to login page if user is not logged in and trying to access a restricted page
  const publicPages = ["/", "/page2", "/login"];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem("user");

  if (authRequired && !loggedIn) {
    console.log("route to Login from " + to.path);
    return next({ name: "Login" });
  }
  next();
});
