#!/bin/sh

# Get filename, which will also upload the file
RETURNED_PATH=$(./run-upload.sh)

# Create some variables
BASE_FILE=$(basename $RETURNED_PATH | sed "s/\(.*\)\.pdf/\1/")
TEX_FILE="$BASE_FILE.tex"
PDF_FILE="$BASE_FILE.pdf"
U_DIR="uploads"
P_DIR="pdf"
D_DIR="downloads"
U_PATH="$U_DIR/$TEX_FILE"
P_PATH="$P_DIR/$PDF_FILE"
D_PATH="$D_DIR/$PDF_FILE"

# Show that it was uploaded
echo
echo File uploaded?
ls "$U_PATH"
if [ "$?" -ne "0" ]; then
  echo Not uploaded... Exiting
  exit
else
  echo Success
fi

# Show that it was converted after it got uploaded
echo
echo File converted?
ls "$P_PATH"
if [ "$?" -ne "0" ]; then
  echo Not converted... Exiting
  exit
else
  echo Success
fi

# Show that it was downloaded
echo
wget "http://localhost:8080/pdf/$PDF_FILE" -P "$D_DIR"
echo File downloaded?
ls "$D_PATH"
if [ "$?" -ne "0" ]; then
  echo Not downloaded... Exiting
  exit
else
  echo Success
fi
