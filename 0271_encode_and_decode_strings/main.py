import unittest


def encode(strs):
    # write your code here
    out = ""
    for s in strs:
        out += str(len(s)) + '#' + s

    return out


def decode(s):
    # write your code here
    out = []
    i = 0
    while i < len(s):
        j = i
        while s[j] != '#':
            j += 1
        num = int(s[i:j])
        out.append(s[j + 1:j + 1 + num])
        i = j + 1 + num

    return out


class TestEncodeDecode(unittest.TestCase):
    def test_encode_decode(self):
        self.assertEqual(
            encode(["lint", "code", "love", "you"]), "4#lint4#code4#love3#you")
        self.assertEqual(
            decode("4#lint4#code4#love3#you"), ["lint", "code", "love", "you"])


if __name__ == '__main__':
    unittest.main()
