// +build darwin

package system

import (
	"crypto/sha1"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pborman/uuid"

	pb_info "github.com/mickep76/grpc-exec-example/info"
)

func Get() (*pb_info.System, error) {
	s := &pb_info.System{
		Manufacturer: "Apple Inc.",
		Kernel:       "Darwin",
		Os:           "Mac OS X",
	}

	s.Hostname, _ = os.Hostname()

	o, err := exec.Command("system_profiler", "SPHardwareDataType").Output()
	if err != nil {
		return nil, err
	}

	for _, l := range strings.Split(string(o), "\n") {
		a := strings.Split(l, ":")
		if len(a) < 2 {
			continue
		}

		k := strings.TrimSpace(a[0])
		v := strings.TrimSpace(a[1])
		switch k {
		case "Model Name":
			s.Product = v
		case "Model Identifier":
			s.ProductVersion = v
		case "Serial Number (system)":
			s.SerialNumber = v
		case "Boot ROM Version":
			s.BootRomVersion = v
		case "SMC Version (system)":
			s.SmcVersion = v
		}
	}

	o, err = exec.Command("sysctl", "-a").Output()
	if err != nil {
		return nil, err
	}

	for _, l := range strings.Split(string(o), "\n") {
		a := strings.SplitN(l, ":", 2)
		if len(a) < 2 {
			continue
		}

		k := strings.TrimSpace(a[0])
		v := strings.TrimSpace(a[1])
		switch k {
		case "hw.memsize":
			i, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, err
			}
			s.MemoryB = i
			s.MemoryGb = uint32(i / 1024 / 1024 / 1024)
		case "machdep.cpu.core_count":
			i, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			s.CpuCoresPerSocket = uint32(i)
		case "hw.physicalcpu_max":
			i, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			s.CpuPhysicalCores = uint32(i)
		case "hw.logicalcpu_max":
			i, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			s.CpuLogicalCores = uint32(i)
		case "machdep.cpu.brand_string":
			s.CpuModel = v
		case "machdep.cpu.features":
			s.CpuFlags = v
		case "kern.osproductversion":
			s.OsVersion = v
		case "kern.osversion":
			s.OsBuild = v
		case "kern.version":
			s.KernelVersion = v
		case "kern.osrelease":
			s.KernelRelease = v
		}
	}

	s.CpuSockets = s.CpuPhysicalCores / s.CpuCoresPerSocket
	s.CpuThreadsPerCore = s.CpuLogicalCores / s.CpuSockets / s.CpuCoresPerSocket

	s.Uuid = uuid.NewHash(sha1.New(), uuid.UUID{}, []byte(s.Hostname+s.Product+s.ProductVersion+s.SerialNumber), 5).String()

	return s, nil
}
