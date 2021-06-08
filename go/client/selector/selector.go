package selector

import (
	"errors"
	"hash/crc32"
	"net"
	"strings"
	"sync"
	"time"
)

// All is a Selector that returns all servers
type All struct {
	sync.RWMutex
	servers []string
}

// Shard is a Selector that shards to a single server
type Shard struct {
	sync.RWMutex
	servers []string
}

func (sa *All) Get(topic string) ([]string, error) {
	sa.RLock()
	if len(sa.servers) == 0 {
		sa.RUnlock()
		return nil, errors.New("no servers")
	}
	var isLiveServers []string
	for _, server := range sa.servers {
		if strings.Contains(server, "//") {
			if sa.judgeLive(strings.Split(server, "//")[1]) {
				isLiveServers = append(isLiveServers, server)
			}
		} else {
			if sa.judgeLive(server) {
				isLiveServers = append(isLiveServers, server)
			}
		}
	}
	servers := isLiveServers
	sa.RUnlock()
	return servers, nil
}

func (sa *All) Set(servers ...string) error {
	sa.Lock()
	sa.servers = servers
	sa.Unlock()
	return nil
}

func (sa *All) judgeLive(server string) bool {
	conn, err := net.DialTimeout("tcp", server, 3*time.Second)
	if err != nil {
		return false
	}
	if conn != nil {
		return true
	}
	return false
}

func (ss *Shard) Get(topic string) ([]string, error) {
	ss.RLock()
	length := len(ss.servers)
	if length == 0 {
		ss.RUnlock()
		return nil, errors.New("no servers")
	}
	if length == 1 {
		servers := ss.servers
		ss.RUnlock()
		return servers, nil
	}
	cs := crc32.ChecksumIEEE([]byte(topic))
	server := ss.servers[cs%uint32(length)]
	ss.RUnlock()
	return []string{server}, nil
}

func (ss *Shard) Set(servers ...string) error {
	ss.Lock()
	ss.servers = servers
	ss.Unlock()
	return nil
}
