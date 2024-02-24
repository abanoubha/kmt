# kmt

A software to create webpages dynamically, and serve them statically.

This is an ongoing project to achieve this [VISION](VISION.md). And Here are some [resources and references of old plans and others' projects](res-refs.md).

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

## Commands

```sh
# clone the project into a directory called 'kmt'
git clone --depth 1 --recursive -b main https://github.com/abanoubha/kmt.git

# download all the libs/modules
go mod tidy

# run the project
go run src/*.go

# build the project, produce a binary/executable
go build -o kmt src/*.go

# run the executable/binary
./kmt

# run tests
go test
```

## Markdown frontmatter

```md
---
title: "title of a blog post about Linux"
date: 2021-10-13T09:45:00+02:00
draft: false
category: "Linux"
---
the body of the article or blog post.
```

The link will be `/linux/some-file-name.html`. The blueprint is `/category-name/file-name.html`.

## Tasks

- [x] (VISION) no database, just plain text files such as markdown files
- [ ] create webpages (blog posts) dynamically via web interface
- [ ] create webpages (blog posts) dynamically by editing the markdown file
- [ ] edit webpages (blog posts) dynamically via web interface
- [ ] edit webpages (blog posts) dynamically by editing the markdown file
- [ ] generate static HTML webpages with CSS/JS, then save them in `/public` dir to serve them to visitors
- [ ] web interface let you edit/create webpage in chunks like gutenberg
- [ ] use [spcss](https://github.com/susam/spcss) by default
- [ ] use [tailwindcss](https://tailwindcss.com/) for advanced theme(s)
- [ ] SEO meta tags
- [ ] robots.txt
- [ ] add canonical url for each page
- [ ] generate sitemap after create/edit
- [ ] RSS XML feed
- [ ] search page, indexing content, 'search' feature
- [ ] notify MS Bing SE that sitemap changed after create/edit
- [ ] notify Google SE that sitemap changed after create/edit
- [ ] widget: share buttons
- [ ] [text to speech](https://developer.mozilla.org/en-US/docs/Web/API/Web_Speech_API)
- [ ] [use local font if available]<https://developer.mozilla.org/en-US/docs/Web/API/Local_Font_Access_API>
- [ ] multilanguage support (i18n)
- [ ] support auto dark mode
- [ ] Auth to create/edit
- [ ] lossless compression for images
- [ ] resize huge images to fit the size needed for. (_hard resize_)
- [ ] breadcrumb
- [ ] `loading="lazy"` for `<img>` and `<iframe>`
- [ ] use __picture__ with __WebP__, __PNG__ and __JPEG__
- [ ] redirects
- [ ] embed YouTube videos
- [ ] embed iframes
- [ ] embed tweets
