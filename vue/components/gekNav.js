Vue.component('gek-nav', {
  template: 
  /*html*/
  `            <!-- Navigation -->
  <div class="bg-white">
      <div class="content py-3">
          <!-- Main Navigation -->
          <div id="main-navigation" class="d-none d-lg-block mt-2 mt-lg-0">
              <ul class="nav-main nav-main-horizontal nav-main-hover">
                  <li class="nav-main-item">
                      <a class="nav-main-link nav-main-link-submenu" data-toggle="submenu"
                          aria-haspopup="true" aria-expanded="false" href="#">
                          <i class="nav-main-link-icon si si-speedometer"></i>
                          <span
                              class="nav-main-link-name">Benutzerfunktionen</span>
                      </a>
                      <ul class="nav-main-submenu">
                          <li class="nav-main-item">
                                <span class="nav-main-link-name"><router-link to="/page1" class="nav-main-link">Seite1</router-link></span>
                          </li>
                      </ul>
                  </li>
                  <li class="nav-main-heading">Heading</li>
                  <li class="nav-main-item">
                      <a class="nav-main-link nav-main-link-submenu" data-toggle="submenu"
                          aria-haspopup="true" aria-expanded="false" href="#">
                          <i class="nav-main-link-icon si si-settings"></i>
                          <span
                              class="nav-main-link-name">Administration</span>
                      </a>
                      <ul class="nav-main-submenu">
                          <li class="nav-main-item">
                              <a class="nav-main-link" href="/entitylistuser">
                                  <span
                                      class="nav-main-link-name">Benutzer</span>
                              </a>
                          </li>
                          <li class="nav-main-item">
                              <a class="nav-main-link" href="/entitylistcontact">
                                  <span
                                      class="nav-main-link-name">Kontakt</span>
                              </a>
                          </li>
                      </ul>
                  </li>
              </ul>
          </div>
          <!-- END Main Navigation -->
      </div>
  </div>
  <!-- END Navigation -->
`,
  data() {
    return {
        product: 'Socks',
        brand: 'Vue Mastery',
        selectedVariant: 0,
        details: ['50% cotton', '30% wool', '20% polyester'],
        variants: [
          { id: 2234, color: 'green', image: './assets/images/socks_green.jpg', quantity: 50 },
          { id: 2235, color: 'blue', image: './assets/images/socks_blue.jpg', quantity: 0 },
        ],
        reviews: []
    }
  },
  methods: {
      addToCart() {
          this.$emit('add-to-cart', this.variants[this.selectedVariant].id)
      },
      updateVariant(index) {
          this.selectedVariant = index
      },
      addReview(review) {
        this.reviews.push(review)
      }
  },
  computed: {
    user () {
      return this.$store.state.user
    }
  },
})