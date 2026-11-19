"""Typeset the Markdown manuscript as a print-sized PDF."""

import argparse
from html import escape
from pathlib import Path
import re

from reportlab.lib import colors
from reportlab.lib.styles import ParagraphStyle
from reportlab.lib.units import inch
from reportlab.pdfbase.pdfmetrics import stringWidth
from reportlab.platypus import (
    BaseDocTemplate,
    Frame,
    KeepTogether,
    NextPageTemplate,
    PageBreak,
    PageTemplate,
    Paragraph,
    Preformatted,
    Spacer,
)
from reportlab.platypus.tableofcontents import TableOfContents


ROOT = Path(__file__).resolve().parents[1]
WIDTH, HEIGHT = 6 * inch, 9 * inch
INK = colors.HexColor("#17232C")
ACCENT = colors.HexColor("#096C66")
PAPER = colors.HexColor("#EDF5F3")
SIDE = 45

BODY = ParagraphStyle(
    "Body", fontName="Times-Roman", fontSize=10.8, leading=13.5,
    textColor=INK, spaceAfter=7, allowWidows=0, allowOrphans=0,
)
BULLET = ParagraphStyle("Bullet", parent=BODY, leftIndent=12, bulletIndent=1)
CHAPTER = ParagraphStyle(
    "Chapter", fontName="Helvetica-Bold", fontSize=23, leading=27,
    textColor=INK, spaceBefore=10, spaceAfter=17, keepWithNext=True,
)
SECTION = ParagraphStyle(
    "Section", fontName="Helvetica-Bold", fontSize=13.4, leading=17,
    textColor=ACCENT, spaceBefore=12, spaceAfter=6, keepWithNext=True,
)
CODE = ParagraphStyle(
    "Code", fontName="Courier", fontSize=8, leading=10.4, textColor=INK,
    borderColor=PAPER, borderWidth=1, backColor=PAPER,
    borderPadding=7, leftIndent=8, rightIndent=8, spaceBefore=4, spaceAfter=13,
)
CONTENTS = ParagraphStyle(
    "Contents", fontName="Helvetica", fontSize=10, leading=23,
    textColor=INK, leftIndent=4,
)
INLINE = re.compile(r"(`[^`]+`|\*\*[^*]+\*\*|\[[^\]]+\]\(https?://[^)]+\))")


def inline(text):
    parts = INLINE.split(text)
    for i in range(1, len(parts), 2):
        token = parts[i]
        if token.startswith("`"):
            parts[i] = '<font name="Courier" size="9">' + escape(token[1:-1]) + "</font>"
        elif token.startswith("**"):
            parts[i] = "<b>" + escape(token[2:-2]) + "</b>"
        else:
            label, url = token[1:-1].split("](", 1)
            parts[i] = '<link color="#096C66" href="' + escape(url, quote=True) + '">' + escape(label) + "</link>"
    for i in range(0, len(parts), 2):
        parts[i] = escape(parts[i])
    return "".join(parts)


def chapter(path):
    paragraphs = []
    pending = []
    code = None

    def flush():
        if pending:
            paragraphs.append(Paragraph(inline(" ".join(pending)), BODY))
            pending.clear()

    for line in path.read_text(encoding="utf-8").splitlines():
        if line.startswith("```"):
            flush()
            if code is None:
                code = []
            else:
                lines = [line.expandtabs(4) for line in code]
                if any(stringWidth(line, CODE.fontName, CODE.fontSize) > WIDTH - 2 * SIDE - 28 for line in lines):
                    raise ValueError(f"Code line is too wide in {path}")
                paragraphs.append(KeepTogether([Preformatted("\n".join(lines), CODE)]))
                code = None
        elif code is not None:
            code.append(line)
        elif not line:
            flush()
        elif line.startswith("# "):
            flush()
            heading = Paragraph(escape(line[2:]), CHAPTER)
            heading.anchor = path.stem
            paragraphs.append(heading)
        elif line.startswith("## "):
            flush()
            paragraphs.append(Paragraph(inline(line[3:]), SECTION))
        elif line.startswith("- "):
            flush()
            paragraphs.append(Paragraph(inline(line[2:]), BULLET, bulletText="•"))
        else:
            pending.append(line)
    if code is not None:
        raise ValueError(f"Unclosed code block in {path}")
    flush()
    return paragraphs


def decorate(canvas, doc):
    if doc.page == 1:
        canvas.setFillColor(INK)
        canvas.rect(0, 0, WIDTH, HEIGHT, fill=1, stroke=0)
        canvas.setFillColor(colors.HexColor("#72CEC1"))
        for index, height in enumerate((18, 32, 53, 82, 122)):
            canvas.rect(SIDE + index * 27, 121, 14, height, fill=1, stroke=0)
        canvas.setFont("Helvetica-Bold", 10)
        canvas.drawString(SIDE, 560, "SECOND EDITION")
        canvas.setFillColor(colors.white)
        canvas.setFont("Helvetica-Bold", 51)
        canvas.drawString(SIDE - 2, 473, "Go for")
        canvas.drawString(SIDE - 2, 417, "Rubyists")
        canvas.setFont("Times-Roman", 15)
        canvas.drawString(SIDE, 366, "From your first Go program to a tested HTTP API")
        canvas.setFont("Helvetica", 12)
        canvas.drawString(SIDE, 318, "Muhammadreza Haghiri")
        return

    canvas.setStrokeColor(colors.HexColor("#C5D6D3"))
    canvas.line(SIDE, 42, WIDTH - SIDE, 42)
    canvas.setFillColor(ACCENT)
    canvas.setFont("Helvetica", 8)
    canvas.drawString(SIDE, 27, "GO FOR RUBYISTS")
    canvas.drawRightString(WIDTH - SIDE, 27, str(doc.page))


class BookPDF(BaseDocTemplate):
    def afterFlowable(self, flowable):
        anchor = getattr(flowable, "anchor", None)
        if anchor:
            title = flowable.getPlainText()
            self.canv.bookmarkPage(anchor)
            self.canv.addOutlineEntry(title, anchor)
            self.notify("TOCEntry", (0, title, self.page, anchor))


def build(output):
    paths = re.findall(r"^\s*- \[[^\]]+\]\((book/[^)]+)\)", (ROOT / "SUMMARY.md").read_text(), re.M)
    if len(paths) != len(set(paths)) or set(paths) != {p.relative_to(ROOT).as_posix() for p in (ROOT / "book").glob("*.md")}:
        raise ValueError("The table of contents must list each manuscript file once")
    output.parent.mkdir(parents=True, exist_ok=True)
    doc = BookPDF(
        str(output), pagesize=(WIDTH, HEIGHT), invariant=True,
        title="Go for Rubyists", author="Muhammadreza Haghiri",
        subject="A practical guide to Go for Ruby programmers",
    )
    frame = Frame(SIDE, 57, WIDTH - 2 * SIDE, HEIGHT - 94, leftPadding=0, rightPadding=0)
    doc.addPageTemplates([
        PageTemplate("Book", [frame], onPage=decorate),
        PageTemplate("Blank", [frame]),
    ])
    toc = TableOfContents()
    toc.levelStyles = [CONTENTS]
    toc.dotsMinLevel = 0
    story = [Spacer(1, 1), PageBreak(), Paragraph("Contents", CHAPTER), toc]
    for path in paths:
        contents = chapter(ROOT / path)
        story.extend([PageBreak(), *contents])
        if re.match(r"^\d+\. ", contents[0].getPlainText()):
            story.extend([
                NextPageTemplate("Blank"), PageBreak(),
                Spacer(1, 1), NextPageTemplate("Book"),
            ])
    doc.multiBuild(story)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--output", type=Path, default=ROOT / "compiled/go-for-rubyists.pdf")
    build(parser.parse_args().output)
