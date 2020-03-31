# HowTo
This is the place to describe HowTo's
# Icons with SVG
webskeleton uses an svg iconset originally served from [Bytesize Icons] (https://github.com/danklammer/bytesize-icons)

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
# Embed resources (templates, css, js, ...)
To serve all resources from inside the executable webskeleton uses [packr2] (https://github.com/gobuffalo/packr/tree/master/v2)

Before building the project you have to generate a resource file with:
`packr2`

After building the project you should clean the generated packr file with:
`packr2 clean`

When no packr resource file is available inside the project dir, then packr searches for the resource files on disk. This should be the default behavior during development.
