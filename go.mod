module github.com/bluenviron/mediamtx

go 1.22

require (
	code.cloudfoundry.org/bytefmt v0.0.0
	github.com/MicahParks/jwkset v0.5.17
	github.com/MicahParks/keyfunc/v3 v3.3.2
	github.com/abema/go-mp4 v1.2.0
	github.com/alecthomas/kong v0.9.0
	github.com/bluenviron/gohlslib v1.3.2
	github.com/bluenviron/gortsplib/v4 v4.9.0
	github.com/bluenviron/mediacommon v1.10.0
	github.com/datarhei/gosrt v0.6.0
	github.com/fsnotify/fsnotify v1.7.0
	github.com/gin-gonic/gin v1.9.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
	github.com/gookit/color v1.5.4
	github.com/gorilla/websocket v1.5.1
	github.com/jackc/pgx/v5 v5.5.3
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/matthewhartstonge/argon2 v1.0.0
	github.com/pion/ice/v2 v2.3.14
	github.com/pion/interceptor v0.1.29
	github.com/pion/logging v0.2.2
	github.com/pion/rtcp v1.2.14
	github.com/pion/rtp v1.8.7-0.20240429002300-bc5124c9d0d0
	github.com/pion/sdp/v3 v3.0.9
	github.com/pion/webrtc/v3 v3.2.28
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.23.0
	golang.org/x/sys v0.20.0
	golang.org/x/term v0.20.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
)

require (
	github.com/asticode/go-astikit v0.42.0 // indirect
	github.com/asticode/go-astits v1.13.0 // indirect
	github.com/benburkert/openpgp v0.0.0-20160410205803-c2471f86866c // indirect
	github.com/bytedance/sonic v1.11.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.18.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.1 // indirect
	github.com/pion/datachannel v1.5.5 // indirect
	github.com/pion/dtls/v2 v2.2.10 // indirect
	github.com/pion/mdns v0.0.12 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/sctp v1.8.12 // indirect
	github.com/pion/srtp/v2 v2.0.18 // indirect
	github.com/pion/stun v0.6.1 // indirect
	github.com/pion/transport/v2 v2.2.4 // indirect
	github.com/pion/turn/v2 v2.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	golang.org/x/arch v0.7.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
  golang.org/x/time v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace code.cloudfoundry.org/bytefmt => github.com/cloudfoundry/bytefmt v0.0.0-20211005130812-5bb3c17173e5

replace github.com/pion/ice/v2 => github.com/aler9/ice/v2 v2.0.0-20231112223552-32d34dfcf3a1

replace github.com/pion/webrtc/v3 => github.com/aler9/webrtc/v3 v3.0.0-20231112223655-e402ed2689c6
