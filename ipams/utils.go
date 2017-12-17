package ipam

import (
	"net"
	"bytes"
	"fmt"
)

type ipVersion int

const (
	v4 = 4
	v6 = 6
)

// It generates the ip address in the passed subnet specified by
// the passed host address ordinal
func GenerateAddress(ordinal uint64, network *net.IPNet) net.IP {
	var address [16]byte

	// Get network portion of IP
	if getAddressVersion(network.IP) == v4 {
		copy(address[:], network.IP.To4())
	} else {
		copy(address[:], network.IP)
	}

	end := len(network.Mask)
	addIntToIP(address[:end], ordinal)

	return net.IP(address[:end])
}

func getAddressVersion(ip net.IP) ipVersion {
	if ip.To4() == nil {
		return v6
	}
	return v4
}

// Adds the ordinal IP to the current array
// 192.168.0.0 + 53 => 192.168.0.53
func addIntToIP(array []byte, ordinal uint64) {
	for i := len(array) - 1; i >= 0; i-- {
		array[i] |= (byte)(ordinal & 0xff)
		ordinal >>= 8
	}
}

var v4inV6MaskPrefix = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

// compareIPMask checks if the passed ip and mask are semantically compatible.
// It returns the byte indexes for the address and mask so that caller can
// do bitwise operations without modifying address representation.
func compareIPMask(ip net.IP, mask net.IPMask) (is int, ms int, err error) {
	// Find the effective starting of address and mask
	if len(ip) == net.IPv6len && ip.To4() != nil {
		is = 12
	}
	if len(ip[is:]) == net.IPv4len && len(mask) == net.IPv6len && bytes.Equal(mask[:12], v4inV6MaskPrefix) {
		ms = 12
	}
	// Check if address and mask are semantically compatible
	if len(ip[is:]) != len(mask[ms:]) {
		err = fmt.Errorf("ip and mask are not compatible: (%#v, %#v)", ip, mask)
	}
	return
}

// GetHostPartIP returns the host portion of the ip address identified by the mask.
// IP address representation is not modified. If address and mask are not compatible
// an error is returned.
func GetHostPartIP(ip net.IP, mask net.IPMask) (net.IP, error) {
	// Find the effective starting of address and mask
	is, ms, err := compareIPMask(ip, mask)
	if err != nil {
		return nil, fmt.Errorf("cannot compute host portion ip address because %s", err)
	}

	// Compute host portion
	out := GetIPCopy(ip)
	for i := 0; i < len(mask[ms:]); i++ {
		out[is+i] &= ^mask[ms+i]
	}

	return out, nil
}

// GetIPCopy returns a copy of the passed IP address
func GetIPCopy(from net.IP) net.IP {
	if from == nil {
		return nil
	}
	to := make(net.IP, len(from))
	copy(to, from)
	return to
}

// Convert an ordinal to the respective IP address
func IpToUint64(ip []byte) (value uint64) {
	cip := GetMinimalIP(ip)
	for i := 0; i < len(cip); i++ {
		j := len(cip) - 1 - i
		value += uint64(cip[i]) << uint(j*8)
	}
	return value
}

// GetMinimalIP returns the address in its shortest form
func GetMinimalIP(ip net.IP) net.IP {
	if ip != nil && ip.To4() != nil {
		return ip.To4()
	}
	return ip
}
