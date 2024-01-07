# kmt

A software to create webpages dynamically, and serve them statically.

- create webpages (blog posts) dynamically.
- edit webpages (blog posts) dynamically.
- creating/editing can be done via __web interface__ or by editing the __markdown file__.
- generate __static HTML webpages__ with CSS/JS, then save them in __/public__ dir to serve them to visitors.
- no database, just plain text files such as markdown files.
- web interface let you edit/create webpage in chunks like gutenberg.
- GIT-based versioning, ability to maintain it in command line via VIM, NeoVIM, ..

_note_: Kmt is named after my country Egypt. Egypt is called Kmt in the ancient egyptian language.

[YouTube Channel](https://www.youtube.com/@abanoubha)
&nbsp; • &nbsp;
[𝕏 (twitter)](https://twitter.com/abanoubha)
&nbsp; • &nbsp;
[LinkedIn](https://www.linkedin.com/in/abanoub-hanna)
&nbsp; • &nbsp;
[Telegram channel](https://t.me/softwarepharaoh)
&nbsp; • &nbsp;
[Facebook page](https://www.facebook.com/SoftwarePharaoh/)

## Default Theme Principles

- UX (user experience)
- minimal
- efficient
- speedy / fast / performant
- SEO (search engine optimization)

## Commands

to create new Go project, run `go mod init github.com/abanoubha/project-name`, then create `main.go` file with `func main(){}` function.

before running the project, run `go mod tidy` then run the project by `go run main.go` or `go run *.go`.

to run tests, run `go test`.

## Features as Tasks

- [ ] use [spcss](https://github.com/susam/spcss) or [tailwindcss](https://tailwindcss.com/)
- [ ] SEO meta tags
- [ ] add canonical url for each page
- [ ] generate sitemap after create/edit
- [ ] rss xml feed
- [ ] indexing content, searching
- [ ] notify MS Bing SE that a URL added
- [ ] notify Google SE that sitemap changed after create/edit
- [ ] widget: share to..
- [ ] feature: text to speech ("read for me" feature). ref:[Web Speech API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Speech_API)
- [ ] use local font if available : <https://developer.mozilla.org/en-US/docs/Web/API/Local_Font_Access_API>
- [ ] translations
- [ ] dark mode | light mode
- [ ] Auth to create/edit
- [ ] lossless compression for images
- [ ] resize huge images to fit the size needed for. (_hard resize_)
- [ ] default theme features: mobile-first, responsive layout
- [ ] default theme features: dark-mode support
- [ ] default theme features: speed page loading time
- [ ] default theme features: SEO optimized ( meta data )
- [ ] default theme features: breadcrumb
- [ ] `loading="lazy"` for `<img>` and `<iframe>`. ([can i use](https://caniuse.com/#feat=loading-lazy-attr)).
- [ ] use __WebP__ image to speed up the page load time. ([can i use](https://caniuse.com/#feat=webp)).
