# web
A software to create webpages dynamically, and serve them statically.

- create webpages (blog posts) dynamically.
- edit webpages (blog posts) dynamically.
- creating/editing can be done via __web interface__ or by editing the __markdown file__.
- generate __static HTML webpages__ with CSS/JS, then save them in __/public__ dir to serve them to visitors.
- no database, just plain text files such as markdown files.

## Ideas & Notes
- Serve cached (generated) static pages –– generated after dynamic pages edited
- lossless compression for images
- resize huge images to fit the size needed for. (_hard resize_)
- read for me (_text-to-speech_)

## Default Theme
The Theme is minimal elegant optimized theme. My vocab are UX (user experience), minimalism, efficiency, speed, performance, optimization, SEO (search engine optimization), 

### Theme Features
- mobile-first, responsive layout
- dark-mode support
- speed page loads
- SEO optimized ( meta data )
- social optimized
- breadcrumb

## Dev Change Log
- v0.1
  - use minimal CSS ([spcss](https://github.com/susam/spcss))
  - `loading="lazy"` for `<img>` and `<iframe>` : almost all modern browsers support it ([can i use](https://caniuse.com/#feat=loading-lazy-attr)).
  - use **WebP** image to speed up the page load time : all modern browsers suppor it ([can i use](https://caniuse.com/#feat=webp)).
  - support dark-mode via CSS (``)

## Features as Tasks

- [ ] SEO meta tags
- [ ] add canonical url for each page
- [ ] generate sitemap after create/edit
- [ ] rss xml feed
- [ ] indexing content, searching
- [ ] notify MS Bing SE that a URL added
- [ ] notify Google SE that sitemap changed after create/edit
- [ ] widget: share to..
- [ ] feature: text to speech ("read for me" feature). ref:[Web Speech API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Speech_API)
- [ ] use local font if available : https://developer.mozilla.org/en-US/docs/Web/API/Local_Font_Access_API
- [ ] translations
- [ ] dark mode | light mode
- [ ] Auth to create/edit

