FROM --platform=${BUILDPLATFORM} golang:1.21-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
ADD . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o steam_discound_tracker_for_discord_${TARGETOS}_${TARGETARCH} main.go

FROM --platform=${BUILDPLATFORM} golang:1.21-alpine

ARG TARGETOS
ARG TARGETARCH
ENV RUN_TARGETOS=$TARGETOS
ENV RUN_TARGETARCH=$TARGETARCH

WORKDIR /app/
COPY --from=builder /app/steam_discound_tracker_for_discord_${TARGETOS}_${TARGETARCH} /app/steam_discound_tracker_for_discord_${TARGETOS}_${TARGETARCH}

ENV WEBHOOK_URL ""
ENV COLOR 15844367
ENV CHECK_CYCLE 30
ENV CURRENCY_SYMBOL "₩"

ENTRYPOINT ["/app/steam_discound_tracker_for_discord_${RUN_TARGETOS}_${RUN_TARGETARCH}","-webhook_url=$WEBHOOK_URL","-color=$COLOR","-check_cycle=$CHECK_CYCLE","-currency_symbol=$CURRENCY_SYMBOL"]