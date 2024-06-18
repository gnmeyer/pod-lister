FROM ubuntu

# Copy the lister binary into the image
COPY /listerLinux /lister


ENTRYPOINT [ "./lister" ]
