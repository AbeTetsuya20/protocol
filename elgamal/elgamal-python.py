import random

# MAX_BIT は鍵の bit 長
MAX_BIT = 100
MAKE_P_Q = False
P_INIT = 243214025379980117728458518153693683424828357373865512470425482311025794530998678691915042056564749285232545430603977729863445075522218259131196386532783120026453503500239803289218265556224244249099926305894370035038658063890151212203457970230373854046430379970545792748775200812334588895724517860420620848527
Q_INIT = 121607012689990058864229259076846841712414178686932756235212741155512897265499339345957521028282374642616272715301988864931722537761109129565598193266391560013226751750119901644609132778112122124549963152947185017519329031945075606101728985115186927023215189985272896374387600406167294447862258930210310424263


# n bit の素数を作成する関数
def isPrime(p):
    count = 0

    while count <= 40:
        a = random.randint(1, p)

        num = pow(a, p - 1, p)
        if num == 1:
            count = count + 1
        else:
            return False

    return True


# make_p は n_bit で p を見つける
def make_p(n_bit):
    count = 0
    while True:
        q = random.randint(pow(2, n_bit - 1), pow(2, n_bit) - 1)
        if isPrime(q):
            p = (2 * q) + 1
            # p=2q+1が素数になるかチェック
            if isPrime(p):
                print(f"{count}回目で成功")
                return p, q
            else:
                print(f"{count}回目の失敗")
                count = count + 1


# make_g は 1 < x < (p-1) の g^q mod p == 1 を満たす自然数
def make_g(p, q):
    while True:
        g = random.randint(1, p - 1)
        if pow(g, q, p) == 1:
            return g


def make_random(q):
    return random.randint(1, q - 1)


def make_y(g, x, p):
    return pow(g, x, p)


def Enc(p, y, g, r, m):
    c1 = pow(g, r, p)
    c2 = pow(y, r, p)
    c = (c1, (m * c2) % p)
    return c


def Dec(x, c, p):
    c1 = c[0]
    c2 = c[1]
    m = c2 * pow(c1, -x, p)
    m = m % p
    return m


def message_printer(message, message_enc, message_enc_dec):
    print("平文   : ", message)
    print("暗号化後: ", message_enc)
    print("復号化後: ", message_enc_dec)


def key_printer(key):
    print("--- 秘密鍵 ---")
    print("x: ", key["x"])
    print("--- 公開鍵 ---")
    print("p: ", key["p"])
    print("q: ", key["q"])
    print("y: ", key["y"])
    print("--- その他 ---")
    print("r: ", key["r"])
    print("g: ", key["g"])


def key_test(key):
    print("-----テストを開始します-----")
    x = key["x"]
    p = key["p"]
    q = key["q"]
    y = key["y"]
    r = key["r"]
    g = key["g"]

    try:
        if not isPrime(p):
            raise ValueError("p error: p の値が素数ではない")
        if not isPrime(q):
            raise ValueError("q error: q の値が素数ではない")
        if not (p == 2*q + 1):
            raise ValueError("q error: p = 2q + 1 を満たさない")
        if y != pow(g, x, p):
            raise ValueError("y error: y = g^x mod p を満たさない")

        print("All Test Passed!")
        print("")
    except ValueError as e:
        print(e)
        exit(1)



def main():
    global MAX_BIT, MAKE_P_Q, P_INIT, Q_INIT

    print("Start Elgamal Chipper")

    key = {}

    # 素数 p, q の生成には時間がかかるため、通常は初期値を用いる
    if MAKE_P_Q:
        key["p"], key["q"] = make_p(MAX_BIT)
    else:
        key["p"] = P_INIT
        key["q"] = Q_INIT

    key["g"] = make_g(key["p"], key["q"])
    key["x"] = make_random(key["q"])
    key["y"] = make_y(key["g"], key["x"], key["p"])
    key["r"] = make_random(key["q"])

    key_printer(key)
    key_test(key)

    m = 100
    m_enc = Enc(key["p"], key["y"], key["g"], key["r"], m)
    m_enc_dec = Dec(key["x"], m_enc, key["p"])

    message_printer(m, m_enc, m_enc_dec)


if __name__ == "__main__":
    print("実行開始")
    main()
