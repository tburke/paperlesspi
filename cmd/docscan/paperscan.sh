for i in scan_*.pnm; do mogrify -shave 50x5 "${i}"; done
for f in ./*.pnm; do unpaper --size "a4" --overwrite "$f" `echo "$f" | sed 's/scan/scan_unpaper/g'`; rm -f "$f"; done
for i in scan_*.pnm; do echo "${i}"; convert "${i}" -contrast-stretch 1% -level 29%,76% "${i}.tif"; done
for i in scan_*.pnm.tif; do echo "${i}"; tesseract "$i" "$i" -l eng pdf; done
export FILE_NAME=scan_`date +%Y%m%d-%H%M%S`
pdftk *.tif.pdf cat output "$FILE_NAME.pdf"

