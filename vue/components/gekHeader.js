Vue.component('gek-header', {
  props: {
    mainHeader: {
      type: String,
      required: true
    },
    user: {
      type: String,
      default: "maxiUser"
    }
  },
  template: 
  /*html*/
  `        <!-- Header -->
  <header id="page-header">
      <!-- Header Content -->
      <div class="content-header">
          <!-- Left Section -->
          <div class="d-flex align-items-center">
              <!-- Logo -->
                <span class="font-w700 font-size-h5 text-dual"><router-link to="/home" class="font-w700 font-size-h5">{{ mainHeader }}</router-link></span>
              <!-- END Logo -->

          </div>
          <!-- END Left Section -->

          <!-- Right Section -->
          <div class="d-flex align-items-center">
              <!-- User Dropdown -->
              <div class="dropdown d-inline-block ml-2">
                  <button type="button" class="btn btn-sm btn-dual" id="page-header-user-dropdown"
                      data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                      <img class="rounded" src="/assets/media/avatars/avatar10.jpg" alt="Header Avatar"
                          style="width: 18px;">
                      <span class="d-none d-sm-inline-block ml-1">{{ user }}</span>
                      <i class="fa fa-fw fa-angle-down d-none d-sm-inline-block"></i>
                  </button>
                  <div class="dropdown-menu dropdown-menu-right p-0 border-0 font-size-sm"
                      aria-labelledby="page-header-user-dropdown">
                      <div class="p-3 text-center bg-primary">
                          <img class="img-avatar img-avatar48 img-avatar-thumb"
                              src="/assets/media/avatars/avatar10.jpg" alt="">
                      </div>
                      <div class="p-2">
                          <h5 class="dropdown-header text-uppercase">Aktionen</h5>
                          <a class="dropdown-item d-flex align-items-center justify-content-between"
                              href="/logout">
                              <span>Logout</span>
                              <i class="si si-logout ml-1"></i>
                          </a>
                      </div>
                  </div>
              </div>
              <!-- END User Dropdown -->
          </div>
          <!-- END Right Section -->
      </div>
      <!-- END Header Content -->

      <!-- Header Loader -->
      <!-- Please check out the Loaders page under Components category to see examples of showing/hiding it -->
      <div id="page-header-loader" class="overlay-header bg-primary-lighter">
          <div class="content-header">
              <div class="w-100 text-center">
                  <i class="fa fa-fw fa-circle-notch fa-spin text-primary"></i>
              </div>
          </div>
      </div>
      <!-- END Header Loader -->
  </header>
  <!-- END Header -->
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
      title() {
          return this.brand + ' ' + this.product
      },
      image() {
          return this.variants[this.selectedVariant].image
      },
      inStock() {
          return this.variants[this.selectedVariant].quantity
      },
      shipping() {
        if (this.premium) {
          return 'Free'
        }
        return 2.99
      }
  }
})