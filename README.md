# stepperphotogopi
Control a stepper motor and a photo machine to make 360 photos to an object, using a raspberry pi programmed in go, all with a web interface
## Scope of this project
This is a little project I am creating for rotating an object 360 degree and make photos of it every x degrees. Then, it should present the user these photos so that they can be used in an online shop. My raspberry actually is in a private network so right by now no authentication will be done

## About webp

I used webp as the image format. So I translate from jpeg to webp in this app. Used go-webpin ( github.com/CapsLock-Studio/go-webpbin ) for encoding, but
also needed to install libwebp. For arm, you should:

apt install sudo apt install libjpeg-dev autoconf automake make gcc g++ wget

wget https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.2.1.tar.gz && \
tar -xvzf libwebp-1.2.1.tar.gz && \
mv libwebp-1.2.1 libwebp && \
rm libwebp-1.2.1.tar.gz && \
cd /libwebp && \
./configure && \
make && \
make install \
cd .. \
rm -rf libwebp

Then add the webp libraries to ldconfig

sudo ldconfig /usr/local/lib/

Finally, just export the three env variables:
export SKIP_DOWNLOAD=true
export VENDOR_PATH=/usr/local/bin
export LIBWEBP_VERSION=1.2.1
