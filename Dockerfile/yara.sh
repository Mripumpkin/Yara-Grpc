# yum install -y automake libtool make gcc pkg-config
 
# wget https://github.91chifun.workers.dev/VirusTotal/yara/archive/refs/tags/v4.1.1.tar.gz
 
# tar -zxvf v4.1.1.tar.gz  
# cd "/tmp/yara-4.2.2/"
# ./bootstrap.sh 
# ./configure --disable-shared --enable-static --without-crypto
# make && make install
# cp "/usr/local/lib/pkgconfig/yara.pc"   "/usr/lib64/pkgconfig" 
 
export YARA_SRC="/tmp/yara-4.2.2"
export CGO_CFLAGS="-I${YARA_SRC}/libyara/include"
export CGO_LDFLAGS="-L${YARA_SRC}/libyara/.libs -lyara -lm"

echo "------------------------end---------------------------------------" 
# go get github.com/hillu/go-yara/v4