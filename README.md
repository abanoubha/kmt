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
go run *.go

# build the project, produce a binary/executable
go build -o kmt .

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

## Roadmap: tasks & to-do

- v0.1.0
  - (VISION) no database, just plain text files such as markdown files
- v0.2.0
  - WIP
- next
  - create a post via web interface (localhost or actual server/hosting)
  - create webpages (blog posts) dynamically by editing the markdown file
  - edit webpages (blog posts) dynamically via web interface
  - edit webpages (blog posts) dynamically by editing the markdown file
  - generate static HTML webpages with CSS/JS, then save them in `/public` dir to serve them to visitors
  - web interface let you edit/create webpage in chunks like gutenberg
  - use [spcss](https://github.com/susam/spcss) by default
  - use [tailwindcss](https://tailwindcss.com/) for advanced theme(s)
  - SEO meta tags
  - robots.txt
  - add canonical url for each page
  - generate sitemap after create/edit
  - RSS XML feed
  - search page, indexing content, 'search' feature
  - notify MS Bing SE that sitemap changed after create/edit
  - notify Google SE that sitemap changed after create/edit
  - widget: share buttons
  - [use local font if available](https://developer.mozilla.org/en-US/docs/Web/API/Local_Font_Access_API)
  - multilanguage support (i18n)
  - support auto dark mode
  - Auth to create/edit
  - lossless compression for images
  - resize huge images to fit the size needed for. (_hard resize_)
  - breadcrumb
  - `loading="lazy"` for `<img>` and `<iframe>`
  - use __picture__ with __WebP__, __PNG__ and __JPEG__
  - redirects
  - embed YouTube videos
  - embed iframes
  - embed tweets
  - tool: OCR for English & Arabic
  - tool: convert numbers to & from Arabic, English and Persian
  - tool: QR Code generator
  - tool: QR code reader
  - tool: image sharing service (مشاركة الصور)
  - tool: image format converter (محول صيغ الصور)
  - tool: image compressor/optimizer (ضاغط حجم الصور)
  - tool: words counter (عداد الكلمات)
  - tool: secure password generator (مولد كلمات المرور)
  - tool: Markdown Editor (محرر مارك داون)
  - tool: image &rarr; Base64
  - tool: Base64 &rarr; image
  - tool: icon maker (صانع الأيقونات)
  - tool: show my IP (معرفة IP الخاص بي)
  - tool: WordPress theme inspector (مكتشف قوالب ووردبريس)
  - tool: domain info (معلومات حول النطاق / الدومين)
  - tool: url encoder (تشفير الرابط)
  - tool: url decoder (فك تشفير الرابط)
  - tool: MD5 generator/calculator (مولد MD5)
  - tool: color converter (محول الألوان)
  - tool: HTML minifier (مصغّر HTML)
  - tool: HTML prettifier (مجمل HTML)
  - tool: HTML encoder (مشفر HTML)
  - tool: HTML decoder (فك تشفير HTML)
  - tool: CSS minifier (مصغر CSS)
  - tool: CSS prettifier (مجمل CSS)
  - tool: Javascript minifier (مصغر Javascript)
  - tool: Javascript prettifier (مجمل Javascript)
  - tool: binary to text (ثنائي إلى نص)
  - tool: text to binary (نص إلى ثنائي)
  - tool: YouTube video thumbnail downloader (تحميل الصورة المصغرة الخاصة بـ فيديو على يوتيوب)
