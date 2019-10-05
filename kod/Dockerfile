############################
# STEP 1 build executable binary
############################
FROM golang:latest as Builder
# Install git.
# Git is required for fetching the dependencies.
# RUN apk update && apk add --no-cache git
# Create appuser.
# RUN adduser -D -g '' appuser
RUN mkdir /kod
copy . /kod/
WORKDIR /kod
# COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d 
# Using go mod.
# RUN go mod download
# RUN go mod verify
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
############################
# STEP 2 build a small image
############################
FROM scratch
# Import the user and group files from the builder.
# COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable.   
copy --from=Builder /kod/main . 

# Use an unprivileged user.
# USER appuser
# Run the hello binary.
CMD ["./main"]