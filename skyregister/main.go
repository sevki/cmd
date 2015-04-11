package main

import (
	"encoding/json"
	"flag"
	"net"
	"os"
	"fmt"
	"log"

	etcd "github.com/coreos/go-etcd/etcd"

	skymsg "github.com/skynetservices/skydns/msg"
)

var (
	etcdClient = flag.String("etcd", "http://0.0.0.0:4001", "etcd_addr")
	domain     = flag.String("domain", "kubernetes.local", "domain")
)

func addDNS(record, ip string, etcdClient *etcd.Client, domain string) error {
	svc := skymsg.Service{
		Host: ip,
	}

	b, err := json.Marshal(svc)
	if err != nil {
		return err
	}
	// Set with no TTL, and hope that kubernetes events are accurate.

	log.Printf("Setting dns record: %v -> %s:%d\n", record, ip, 27017)
	_, err = etcdClient.RawSet(skymsg.Path(fmt.Spritnf("%s.%s", record, domain)), string(b), uint64(0))
	return err
}

func main() {
	flag.Parse()
	client := etcd.NewClient([]string{*etcdClient})
	if client == nil {
		panic(nil)
	}
	client.SyncCluster()

	addrs, _ := net.InterfaceAddrs()

	for _, v := range addrs {
		this, _, _ := net.ParseCIDR(v.String())

		//		fmt.Println(this.Equal(home))
		if this.IsLinkLocalUnicast() || this.IsLoopback() {
			continue
		}

		hostname, _ := os.Hostname()
		addDNS(hostname, this.String(), client, *domain)
	}
	client.SyncCluster()

}
