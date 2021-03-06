// +build linux

package system

import (
	"crypto/sha1"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pborman/uuid"

	pb_info "github.com/mickep76/grpc-exec-example/info"
)

func readFile(fn string) (string, error) {
	b, err := ioutil.ReadFile(fn)
	return strings.TrimSpace(strings.TrimRight(string(b), "\n")), err
}

func Get() (*pb_info.System, error) {
	s := &pb_info.System{
		Kernel: "Linux",
	}

	s.Hostname, _ = os.Hostname()

	var err error
	s.Manufacturer, err = readFile("/sys/devices/virtual/dmi/id/sys_vendor")
	if err != nil {
		return nil, err
	}

	s.Product, err = readFile("/sys/devices/virtual/dmi/id/product_name")
	if err != nil {
		return nil, err
	}

	s.ProductVersion, err = readFile("/sys/devices/virtual/dmi/id/product_version")
	if err != nil {
		return nil, err
	}

	s.SerialNumber, err = readFile("/sys/devices/virtual/dmi/id/product_serial")
	if err != nil {
		return nil, err
	}

	s.BiosVendor, err = readFile("/sys/devices/virtual/dmi/id/bios_vendor")
	if err != nil {
		return nil, err
	}

	s.BiosDate, err = readFile("/sys/devices/virtual/dmi/id/bios_date")
	if err != nil {
		return nil, err
	}

	s.BiosVersion, err = readFile("/sys/devices/virtual/dmi/id/bios_version")
	if err != nil {
		return nil, err
	}

	s.KernelVersion, err = readFile("/proc/sys/kernel/version")
	if err != nil {
		return nil, err
	}

	s.KernelRelease, err = readFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return nil, err
	}

	o, err := readFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	for _, l := range strings.Split(o, "\n") {
		a := strings.SplitN(l, ":", 2)
		if len(a) < 2 {
			continue
		}

		k := strings.TrimSpace(a[0])
		v := strings.TrimSpace(a[1])
		v2 := strings.TrimSpace(strings.Split(v, " ")[0])

		switch k {
		case "MemTotal":
			i, err := strconv.ParseUint(v2, 10, 64)
			if err != nil {
				return nil, err
			}
			s.MemoryB = i * 1024
			s.MemoryGb = uint32(i / 1024 / 1024)
		}
	}

	if err := getCPU(s); err != nil {
		return nil, err
	}

	if err := getOS(s); err != nil {
		return nil, err
	}

	s.Uuid = uuid.NewHash(sha1.New(), uuid.UUID{}, []byte(s.Hostname+s.Product+s.ProductVersion+s.SerialNumber), 5).String()

	return s, nil
}
