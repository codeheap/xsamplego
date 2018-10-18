FROM debian:8

# create the directory where the application will reside
RUN mkdir /authserver

# copy the application files needed for production (in this case, only the binary)
ADD authserver /authserver/authserver

# set the working directory to the application directory
WORKDIR /authserver

# expose the application on port 8001.
# This should be the same as in the port used in the application
EXPOSE 8001

# set the entry point of the container to the application executable
ENTRYPOINT /authserver/authserver