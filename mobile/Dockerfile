FROM cirrusci/flutter

COPY . /app
WORKDIR /app

USER root

RUN rm -f .packages

RUN flutter pub get
RUN flutter clean
RUN flutter build apk  --split-per-abi --no-sound-null-safety

RUN mkdir /mobile_app/
RUN cp build/app/outputs/flutter-apk/app.apk /mobile_app/client.apk
