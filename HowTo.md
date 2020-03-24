# HowTo
This is the place to describe HowTo's
# Icons with SVG
webskeleton uses an svg iconset originally served from Bytesize Icons (https://github.com/danklammer/bytesize-icons)

All svg icons are defined in the file gks-icons.svg, styled with css rule `gks-icon` and integrated in your html with links.

Example to show the edit icon i-edit:

` <svg class="gks-icon">
    <use xlink:href="/assets/svg/gks-icons.svg#i-edit"></use>
  </svg>
`

Example, using a link to start editing after clicking the icon:

` <a href="#" data-toggle="modal" data-target="#userEditModal">
    <svg class="gks-icon">
      <use xlink:href="/assets/svg/gks-icons.svg#i-edit"></use>
    </svg>
  </a>
`