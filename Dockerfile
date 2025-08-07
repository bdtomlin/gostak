FROM golang:1.24.6
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/bokwoon95/wgo@latest

# Copy all project source to working directory
WORKDIR /app
COPY . .
RUN chmod +x /app/entrypoint.sh
