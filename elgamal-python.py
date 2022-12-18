import random

# MAX_BIT は鍵の bit 長
MAX_BIT = 100


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
    p = 0
    q = 0

    while True:
        q = random.randint(pow(2, n_bit - 1), pow(2, n_bit) - 1)
        if isPrime(q):
            p = (2 * q) + 1
            # p=2q+1が素数になるかチェック
            if isPrime(p):
                print(f"{count}回目で成功")
                break
            else:
                print(f"{count}回目の失敗")
                count = count + 1

    return p, q


def main():
    print()


if __name__ == "__main__":
    print("実行開始")
    main()
