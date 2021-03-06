package support

import "strings"

var mapMimeExtension = map[string]string{
	"application/x-xz":            ".xz",
	"application/gzip":            ".gz",
	"application/x-7z-compressed": ".7z",
	"application/zip":             ".zip",
	"application/x-tar":           ".tar",
	"application/x-xar":           ".xar",
	"application/x-bzip2":         ".bz2",
	"application/pdf":             ".pdf",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         ".xlsx",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   ".docx",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx",
	"application/epub+zip":                 ".epub",
	"application/jar":                      ".jar",
	"application/x-ole-storage":            "",
	"application/msword":                   ".doc",
	"application/vnd.ms-powerpoint":        ".ppt",
	"application/vnd.ms-publisher":         ".pub",
	"application/vnd.ms-excel":             ".xls",
	"application/vnd.ms-outlook":           ".msg",
	"application/postscript":               ".ps",
	"application/fits":                     ".fits",
	"application/ogg":                      ".ogg",
	"audio/ogg":                            ".oga",
	"video/ogg":                            ".ogv",
	"text/plain":                           ".txt",
	"text/xml":                             ".xml",
	"application/json":                     ".json",
	"text/csv":                             ".csv",
	"text/tab-separated-values":            ".tsv",
	"application/geo+json":                 ".geojson",
	"application/x-ndjson":                 ".ndjson",
	"text/html":                            ".html",
	"text/x-php":                           ".php",
	"text/rtf":                             ".rtf",
	"application/javascript":               ".js",
	"text/x-lua":                           ".lua",
	"text/x-perl":                          ".pl",
	"application/x-python":                 ".py",
	"text/x-tcl":                           ".tcl",
	"text/vcard":                           ".vcf",
	"text/calendar":                        ".ics",
	"text/markdown":                        ".md",
	"image/svg+xml":                        ".svg",
	"application/rss+xml":                  ".rss",
	"application/atom+xml":                 ".atom",
	"model/x3d+xml":                        ".x3d",
	"application/vnd.google-earth.kml+xml": ".kml",
	"application/x-xliff+xml":              ".xlf",
	"model/vnd.collada+xml":                ".dae",
	"application/gml+xml":                  ".gml",
	"application/gpx+xml":                  ".gpx",
	"application/vnd.garmin.tcx+xml":       ".tcx",
	"application/x-amf":                    ".amf",
	"application/vnd.ms-package.3dmanufacturing-3dmodel+xml": ".3mf",
	"image/png":                        ".png",
	"image/jpeg":                       ".jpg",
	"image/jp2":                        ".jp2",
	"image/jpx":                        ".jpf",
	"image/jpm":                        ".jpm",
	"image/bpg":                        ".bpg",
	"image/gif":                        ".gif",
	"image/webp":                       ".webp",
	"image/tiff":                       ".tiff",
	"image/bmp":                        ".bmp",
	"image/x-icon":                     ".ico",
	"image/x-icns":                     ".icns",
	"image/vnd.adobe.photoshop":        ".psd",
	"image/heic":                       ".heic",
	"image/heic-sequence":              ".heic",
	"image/heif":                       ".heif",
	"image/heif-sequence":              ".heif",
	"audio/mpeg":                       ".mp3",
	"audio/flac":                       ".flac",
	"audio/midi":                       ".midi",
	"audio/ape":                        ".ape",
	"audio/musepack":                   ".mpc",
	"audio/wav":                        ".wav",
	"audio/aiff":                       ".aiff",
	"audio/basic":                      ".au",
	"audio/amr":                        ".amr",
	"audio/aac":                        ".aac",
	"audio/x-unknown":                  ".voc",
	"audio/mp4":                        ".mp4",
	"audio/x-m4a":                      ".m4a",
	"video/x-m4v":                      ".m4v",
	"video/mp4":                        ".mp4",
	"video/webm":                       ".webm",
	"video/mpeg":                       ".mpeg",
	"video/quicktime":                  ".mov",
	"video/3gpp":                       ".3gp",
	"video/3gpp2":                      ".3g2",
	"video/x-msvideo":                  ".avi",
	"video/x-flv":                      ".flv",
	"video/x-matroska":                 ".mkv",
	"video/x-ms-asf":                   ".asf",
	"application/vnd.rn-realmedia-vbr": ".rmvb",
	"application/x-java-applet":        ".class",
	"application/x-shockwave-flash":    ".swf",
	"application/x-chrome-extension":   ".crx",
	"font/ttf":                         ".ttf",
	"font/woff":                        ".woff",
	"font/woff2":                       ".woff2",
	"font/otf":                         ".otf",
	"application/vnd.ms-fontobject":    ".eot",
	"application/wasm":                 ".wasm",
	"application/octet-stream":         ".shp",
	"application/x-dbf":                ".dbf",
	"application/vnd.microsoft.portable-executable":            ".exe",
	"application/x-elf":                                        "",
	"application/x-object":                                     "",
	"application/x-executable":                                 "",
	"application/x-sharedlib":                                  ".so",
	"application/x-coredump":                                   "",
	"application/x-archive":                                    ".a",
	"application/vnd.debian.binary-package":                    ".deb",
	"application/x-rpm":                                        ".rpm",
	"application/dicom":                                        ".dcm",
	"application/vnd.oasis.opendocument.text":                  ".odt",
	"application/vnd.oasis.opendocument.text-template":         ".ott",
	"application/vnd.oasis.opendocument.spreadsheet":           ".ods",
	"application/vnd.oasis.opendocument.spreadsheet-template":  ".ots",
	"application/vnd.oasis.opendocument.presentation":          ".odp",
	"application/vnd.oasis.opendocument.presentation-template": ".otp",
	"application/vnd.oasis.opendocument.graphics":              ".odg",
	"application/vnd.oasis.opendocument.graphics-template":     ".otg",
	"application/vnd.oasis.opendocument.formula":               ".odf",
	"application/x-rar-compressed":                             ".rar",
	"image/vnd.djvu":                                           ".djvu",
	"application/x-mobipocket-ebook":                           ".mobi",
	"application/x-ms-reader":                                  ".lit",
	"application/x-sqlite3":                                    ".sqlite",
	"image/vnd.dwg":                                            ".dwg",
	"application/warc":                                         ".warc",
	"application/vnd.nintendo.snes.rom":                        ".nes",
	"application/x-mach-binary":                                ".macho",
	"audio/qcelp":                                              ".qcp",
	"application/marc":                                         ".mrc",
	"application/x-msaccess":                                   ".mdb",
	"application/zstd":                                         ".zst",
	"application/vnd.ms-cab-compressed":                        ".cab",
	"application/lzip":                                         ".lz",
	"application/x-bittorrent":                                 ".torrent",
	"application/x-cpio":                                       ".cpio",
}

func MimeByExtension(filename string) (string, bool) {
	if filename == "" {
		return "", false
	}
	parts := strings.Split(filename, ".")
	inputExtension := "." + parts[len(parts)-1]
	for mime, extension := range mapMimeExtension {
		if extension == inputExtension {
			return mime, true
		}
	}
	return "", false
}
