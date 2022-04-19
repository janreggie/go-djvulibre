# go-djvulibre

An attempt on porting the [djvulibre](https://github.com/traycold/djvulibre/) library to Go.

Right now I'm just translating the code line-by-line.
I *know* that the code isn't the most idiomatic.
but for now this should work.

I know that Rust should be the ideal language for this,
but I'll port it to *that* language
when I'm done understanding everything about djvulibre,
which is one of the objectives of this port.

## Progress

- 0: No implementation at all
- 1: Implementation written at the first pass
- 2: Implementation written at the second pass and idiomatic per Go guidelines
- X: Implementation may not be needed (subject to change)

### `libdjvu` the library

| File in C++ | Progress | File in Go | Notes |
| --- | --- | --- | --- |
| `Arrays.cpp` | X | slices |
| `BSByteStream.cpp` | 0 | unimplemented |
| `BSEncodeByteStream.cpp` | 0 | unimplemented |
| `ByteStream.cpp` | 0 | umplemented | stdlib io works fine for now |
| `DataPool.cpp` | 0 | `data_pool.go` | idk what this is for (yet) |
| `DjVmDir.cpp` | 0 | unimplemented |
| `DjVmDir0.cpp` | 0 | unimplemented |
| `DjVmDoc.cpp` | 0 | `multidoc.go` |
| `DjVmNav.cpp` | 0 | unimplemented |
| `DjVuAnno.cpp` | 0 | unimplemented |
| `DjVuDocEditor.cpp` | 0 | `editor.go` | fields+methods unimplemented |
| `DjVuDocument.cpp` | 0 | `document.go` | fields+methods unimplemented |
| `DjVuDumpHelper.cpp` | 0 | unimplemented |
| `DjVuErrorList.cpp` | 0 | unimplemented |
| `DjVuFile.cpp` | 0 | `file.go` | everything unimplemented |
| `DjVuFileCache.cpp` | 0 | unimplemented |
| `DjVuGlobal.cpp` | 0 | unimplemented |
| `DjVuGlobalMemory.cpp` | 0 | unimplemented |
| `DjVuImage.cpp` | 0 | unimplemented |
| `DjVuInfo.cpp` | 1 | `info.go` |
| `DjVuMessage.cpp` | 0 | unimplemented |
| `DjVuMessageLite.cpp` | 0 | unimplemented |
| `DjVuNavDir.cpp` | 0 | unimplemented |
| `DjVuPalette.cpp` | 0 | unimplemented |
| `DjVuPort.cpp` | 0 | `port*.go` | idk what this does (yet) |
| `DjVuText.cpp` | 0 | unimplemented |
| `DjVuToPS.cpp` | 0 | unimplemented |
| `GBitmap.cpp` | 0 | `image/bitmap.go` | we're getting there |
| `GContainer.cpp` | 0 | | Thread-safe map,linked-list,array. Use generics. |
| `GException.cpp` | 0 | | This is a custom error type |
| `GIFFManager.cpp` | 0 | unimplemented |
| `GMapAreas.cpp` | 0 | unimplemented |
| `GOS.cpp` | X | stdlib os |
| `GPixmap.cpp` | 0 | `image/pixmap.go` | we're getting there |
| `GRect.cpp` | 1 | `image/rect.go` |
| `GScaler.cpp` | 0 | unimplemented |
| `GSmartPointer.cpp` | X | pointers | Go is a GC language |
| `GString.cpp` | X | stdlib string | strings can store UTF-8 and Go is in UTF-8 |
| `GThreads.cpp` | 0 | | `GMonitor` is a mutex.
| `GURL.cpp` | 0 | `url.go` |
| `GUnicode.cpp` | 0 | | Try stdlib unicode |
| `IFFByteStream.cpp` | 0 | unimplemented |
| `IW44EncodeCodec.cpp` | 0 | unimplemented |
| `IW44Image.cpp` | 0 | unimplemented |
| `JB2EncodeCodec.cpp` | 0 | unimplemented |
| `JB2Image.cpp` | 0 | unimplemented |
| `JPEGDecoder.cpp` | 0 | unimplemented |
| `MMRDecoder.cpp` | 0 | unimplemented |
| `MMX.cpp` | 0 | unimplemented |
| `UnicodeByteStream.cpp` | 0 | unimplemented |
| `XMLParser.cpp` | 0 | unimplemented |
| `XMLTags.cpp` | 0 | unimplemented |
| `ZPCodec.cpp` | 0 | unimplemented |
| `atomic.cpp` | X | stdlib atomic |
| `ddjvuapi.cpp` | 0 | unimplemented |
| `debug.cpp` | 0 | unimplemented | Custom logger? |
| `miniexp.cpp` | 0 | unimplemented | Lisp interpreter |

### `tools` the executables

I haven't implemented any.

## License

GPL v2 or higher.

For the original library:

```none
DjVuLibre-3.5
Copyright (c) 2002  Leon Bottou and Yann Le Cun.
Copyright (c) 2001  AT&T

This software is subject to, and may be distributed under, the
GNU General Public License, either Version 2 of the license,
or (at your option) any later version. The license should have
accompanied the software or you may obtain a copy of the license
from the Free Software Foundation at http://www.fsf.org .

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

DjVuLibre-3.5 is derived from the DjVu(r) Reference Library from
Lizardtech Software.  Lizardtech Software has authorized us to
replace the original DjVu(r) Reference Library notice by the following
text (see doc/lizard2002.djvu and doc/lizardtech2007.djvu):

+------------------------------------------------------------------
| DjVu (r) Reference Library (v. 3.5)
| Copyright (c) 1999-2001 LizardTech, Inc. All Rights Reserved.
| The DjVu Reference Library is protected by U.S. Pat. No.
| 6,058,214 and patents pending.
|
| This software is subject to, and may be distributed under, the
| GNU General Public License, either Version 2 of the license,
| or (at your option) any later version. The license should have
| accompanied the software or you may obtain a copy of the license
| from the Free Software Foundation at http://www.fsf.org .
|
| The computer code originally released by LizardTech under this
| license and unmodified by other parties is deemed "the LIZARDTECH
| ORIGINAL CODE."  Subject to any third party intellectual property
| claims, LizardTech grants recipient a worldwide, royalty-free, 
| non-exclusive license to make, use, sell, or otherwise dispose of 
| the LIZARDTECH ORIGINAL CODE or of programs derived from the 
| LIZARDTECH ORIGINAL CODE in compliance with the terms of the GNU 
| General Public License.   This grant only confers the right to 
| infringe patent claims underlying the LIZARDTECH ORIGINAL CODE to 
| the extent such infringement is reasonably necessary to enable 
| recipient to make, have made, practice, sell, or otherwise dispose 
| of the LIZARDTECH ORIGINAL CODE (or portions thereof) and not to 
| any greater extent that may be necessary to utilize further 
| modifications or combinations.
|
| The LIZARDTECH ORIGINAL CODE is provided "AS IS" WITHOUT WARRANTY
| OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
| TO ANY WARRANTY OF NON-INFRINGEMENT, OR ANY IMPLIED WARRANTY OF
| MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE.
+------------------------------------------------------------------
```
