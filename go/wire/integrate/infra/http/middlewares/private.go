package middlewares

import (
	"net"
	"net/http"
	"strings"
)

var privateIPNet = []net.IPNet{
	net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(8, 32)},
	net.IPNet{IP: net.ParseIP("172.16.0.0"), Mask: net.CIDRMask(12, 32)},
	net.IPNet{IP: net.ParseIP("192.168.0.0"), Mask: net.CIDRMask(16, 32)},
}

func clientIP(r *http.Request) net.IP {
	address := r.Header.Get("X-Real-Ip")
	if address == "" {
		address = r.Header.Get("X-Forwarded-For")
	}
	if address == "" {
		address = strings.Split(r.RemoteAddr, ":")[0]
	}

	return net.ParseIP(address)
}

func isPrivate(ip net.IP) bool {
	private := ip.IsLoopback()

	if !private {
		for _, ipnet := range privateIPNet {
			if ipnet.Contains(ip) {
				private = true
				break
			}
		}
	}

	return private
}

// Private API MiddleWare
func PrivateIPMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPrivate(clientIP(r)) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", 401)
		}
	})
}
