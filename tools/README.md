# Build the PDF

Use Python 3.10+ and ReportLab 4.4.9:

```sh
python3 -m pip install -r tools/requirements.txt
python3 tools/build_pdf.py
```

The build reads the ordered manuscript from `SUMMARY.md`, sets a 6 × 9 inch trim size, and writes `compiled/go-for-rubyists.pdf`. Pass `--output PATH` to write elsewhere. The PDF uses ReportLab's standard print fonts and has a linked table of contents and PDF bookmarks. The renderer supports the Markdown used in the manuscript: paragraphs, headings, unordered lists, links, code spans, bold, and fenced blocks.

Every numbered chapter is followed by a blank page without a running footer.
