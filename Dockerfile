from docker.io/node:22.11.0 AS node-builder
COPY ./frontend /frontend
WORKDIR /frontend
RUN npm ci
RUN npx ng build
FROM docker.io/golang:1.23.4 AS go-builder
COPY ./backend /backend
COPY --from=node-builder /frontend/dist/frontend/browser /backend/cmd/strichliste/frontendDist
WORKDIR /backend
RUN CGO_ENABLED=0 go build -o ./strichliste ./cmd/strichliste/main.go
FROM gcr.io/distroless/static-debian12 AS endstage
COPY --from=go-builder /backend/strichliste /app
EXPOSE 8080
ENTRYPOINT [ "/strichliste" ]
