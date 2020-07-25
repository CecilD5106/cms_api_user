FROM golang:alpine as builder
# Add git to builder container to allow download of packages
RUN apk add git
# Make a directory named build
RUN mkdir build
# Make build the working directory
WORKDIR /build
# Copy the go code into the build directory
COPY . /build/
# Download the go packages
RUN go get -d github.com/gin-gonic/gin
RUN go get -d github.com/go-sql-driver/mysql
# Build the go application and put executable in the build directory
RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main
# Designate alpine as the base container
FROM alpine
# Make a directory named app
RUN mkdir app
# Make app the working directory
WORKDIR /app
# Create a user named appuser to run the application
RUN adduser -S -D -H -h /app appuser
# Set the current user to appuser
USER appuser
# Copy the application from build in the builder container to the app directory
COPY --from=builder /build/main /app/
# Make the container listen on port 8000
EXPOSE 8000
# Desiginate the starting command for the application
ENTRYPOINT ["./main"]