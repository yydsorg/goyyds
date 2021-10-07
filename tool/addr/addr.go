package addr

import (
	"log"
	"net"
)

var (
	privateBlocks []*net.IPNet
)

func init() {
	for _, s := range []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "100.64.0.0/10", "fd00::/8"} {
		cidr, ipNet, err := net.ParseCIDR(s)
		log.Println(cidr, ipNet, err)
		privateBlocks = append(privateBlocks, ipNet)
	}
}

func AppendPrivateBlocks(bs ...string) {
	for _, b := range bs {
		if _, block, err := net.ParseCIDR(b); err == nil {
			log.Println(block)
			privateBlocks = append(privateBlocks, block)
		}
	}
}

func isPrivateIp(n string) bool {
	ip := net.ParseIP(n)
	for _, block := range privateBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func IsLocal(addr string) bool {
	host, _, err := net.SplitHostPort(addr)
	//log.Println(host, port, err)
	if err == nil {
		addr = host
	}

	if addr == "localhost" {
		return true
	}
	for _, ip := range IPs() {
		if addr == ip {
			return true
		}
	}
	return false
}

func IPs() []string {
	ifaces, err := net.Interfaces()

	if err != nil {
		return nil
	}
	var ipAddrs []string

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()

		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil {
				continue
			}
			// dont skip ipv6 addrs
			/*
				ip = ip.To4()
				if ip == nil {
					continue
				}
			*/
			ipAddrs = append(ipAddrs, ip.String())
		}
	}

	return ipAddrs
}
