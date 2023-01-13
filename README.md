
# ディレクトリ構成
## rsa/
- rsa.go, rsa-python.py
  - RSA 暗号の実装が、Go 言語と Python で書かれています。
- rsa_solve.go
  - RSA 暗号を解読する関数が書かれています。
- timer.go
  - RSA の暗号化や、解読する時間を計測する関数が書かれています。

## elgamal/
- elgamal.go, elgamal-python.py
  - Elgamal 暗号の実装が、Go 言語と Python で書かれています。

## main/
### elgamal/
- main.go
  - Elgamal 暗号を実行する main 関数があります。

### rsa/
- main.go
  - RSA 暗号を実行する main 関数があります。

### http/
- main.go
  - HTTP 通信を実行する main 関数があります。

### https/
- main.go
  - HTTPS 通信を実行する main 関数があります。

