# minimal blog system
It is a blogging engine/system/platform that let you create blog posts dynamically (easily), and serve web pages to visitors as static webpages.

The Theme is minimal elegant optimized theme. My vocab are UX (user experience), minimalism, efficiency, speed, performance, optimization, SEO (search engine optimization), 

## features
- supports blogger, wordpress, hugo
- mobile-first, responsive layout
- dark-mode support
- speed page loads
- SEO optimized ( meta data )
- social optimized
- breadcrumb

## Ideas & Notes
- Serve cached (generated) static pages generated when dynamic pages edited
- lossless compression for images
- resize huge images to fit the size needed for. (_hard resize_)
- read for me (_text-to-speech_)

## Dev Change Log
- v0.1
  - use minimal CSS ([spcss](https://github.com/susam/spcss))
  - `loading="lazy"` for `<img>` and `<iframe>` : almost all modern browsers support it ([can i use](https://caniuse.com/#feat=loading-lazy-attr)).
  - use **WebP** image to speed up the page load time : all modern browsers suppor it ([can i use](https://caniuse.com/#feat=webp)).
  - support dark-mode via CSS (``)
